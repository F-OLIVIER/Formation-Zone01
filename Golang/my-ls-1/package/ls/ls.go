package LS

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"

	Annexe "my-ls-1/package/annexe"
	Sort "my-ls-1/package/sort"
	Struct "my-ls-1/package/struct"
)

func Ls(l bool, r bool, R bool, a bool, t bool, arg []string, count int) ([]Struct.FileItem, []string) {
	if len(arg[0]) > 0 {
		if Annexe.CherchePoint(arg[0]) && !a { // Gestion des fichiers caché et de la recursive, return si false
			if !R || len(arg) == 1 {
				return nil, []string{}
			} else {
				return nil, arg[1:]
			}
		}
	}

	if l && count == 0 && !R {
		filePath := "./" + arg[0]
		// Récupère les informations sur le fichier
		fileInfo, err := os.Lstat(filePath)
		// Vérifie si le chemin est un lien symbolique
		if err == nil {
			if fileInfo.Mode()&os.ModeSymlink != 0 {
				linkPath, _ := os.Readlink(filePath)
				_, ftype := Annexe.CheckPermission(fileInfo.Mode(), filePath)
				var toReturn Struct.FileItem
				toReturn.Ftype = ftype
				toReturn.Lastmod = fileInfo.ModTime()
				toReturn.Size = int(fileInfo.Size())
				toReturn.Ftype = string(fileInfo.Mode().String()[0])
				toReturn.Permission = "l" + fileInfo.Mode().String()[1:]
				toReturn.OriginalName = fileInfo.Name()
				toReturn.Name = fileInfo.Name()
				toReturn.Linkname = linkPath
				toReturn.FolderPath = filePath
				if sys := fileInfo.Sys(); sys != nil {
					if stat, ok := sys.(*syscall.Stat_t); ok {
						//recherch major, minor, nlink, nom groupe et nom user
						toReturn.Major, toReturn.Minor = Annexe.DeviceNumber(stat.Rdev)
						toReturn.Link = int(stat.Nlink)
						groupName, _ := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
						toReturn.User = groupName.Name
						userName, _ := user.LookupId(strconv.Itoa(int(stat.Uid)))
						toReturn.Group = userName.Username
					}
				}
				temps := []Struct.FileItem{}
				temps = append(temps, toReturn)
				Annexe.Printlist(temps, l, r, R, a, t)
				if !R || len(arg) == 1 {
					return nil, []string{}
				} else {
					return nil, arg[1:]
				}
			}
		}
	}
	listeFD, errReadDir := os.ReadDir(arg[0])
	folderPath := arg[0]
	if errReadDir != nil { // si on a un fichier à la place d'un dossier, ce bout de code return à la fin
		pwd, _ := os.Getwd()
		fileInfo, err := os.Stat(pwd + "/" + arg[0])
		if err != nil { // si le fichier n'existe pas, on fait rien
			fmt.Println("ls: impossible d'accéder à '" + arg[0] + "': Aucun fichier ou dossier de ce type")
		} else {
			_, ftype := Annexe.CheckPermission(fileInfo.Mode(), pwd+"/"+arg[0])
			temp := []string{pwd + "/" + arg[0]}
			Structures, _ := StructStorage(fileInfo, fileInfo.Mode().String(), ftype, temp)
			temps := []Struct.FileItem{Structures}
			Annexe.Printlist(temps, l, r, R, a, t)
		}
		return nil, []string{}
	}
	var listeFetD []Struct.FileItem
	var perm, ftype string
	// traitement de l'option petit r : affichage des fichiers caché
	// fetch . and .. (gestion dossier . et ..)
	if !r {
		listeFetD = AddDots(arg, folderPath, listeFetD, r)
	}
	for _, file := range listeFD { // Gestion des fichiers dossier autres que . et ..
		temp, _ := file.Info()
		if len(arg) > 0 {
			perm, ftype = Annexe.CheckPermission(temp.Mode(), arg[0]+"/"+file.Name())
		} else {
			perm, ftype = Annexe.CheckPermission(temp.Mode(), "./"+file.Name())
		}
		tempListeFetD, tempArg := StructStorage(temp, perm, ftype, arg)
		arg = tempArg
		tempListeFetD.FolderPath = folderPath // chemin du dossier contenant le fichier
		listeFetD = append(listeFetD, tempListeFetD)
	}
	if r {
		listeFetD = AddDots(arg, folderPath, listeFetD, r)
	}
	if r && !t { // traitement de l'option petit r : tri inversé
		listeFetD = Sort.SortSlice(listeFetD, 1)
	} else if t {
		listeFetD = Sort.SortSliceDate(listeFetD, r)
	} else {
		listeFetD = Sort.SortSlice(listeFetD, 0)
	}
	if !R || len(arg) == 1 {
		return listeFetD, []string{}
	} else {
		return listeFetD, arg[1:]
	}
}

// Fonction qui socke dans la struct de fichier
func StructStorage(file fs.FileInfo, perm string, ftype string, path []string) (Struct.FileItem, []string) {
	var toReturn Struct.FileItem
	toReturn.Lastmod = file.ModTime()
	toReturn.Size = int(file.Size())
	//correction de type de fichier exemple : D = b ; L = l
	//voir https://askubuntu.com/questions/466198/how-do-i-change-the-color-for-directories-with-ls-in-the-console
	if ftype == "D" { //fichier block
		toReturn.Ftype = "b"
		perm = "b" + perm[1:]
		toReturn.Permission = perm
	} else if ftype == "L" { // fichier lien
		toReturn.Ftype = "l"
		perm = "l" + perm[1:]
		toReturn.Permission = perm
	} else if ftype == "t" { // fichier sticky
		toReturn.Ftype = "d"
		perm = "d" + perm[1:9] + "t"
		toReturn.Permission = perm
	} else if ftype == "g" || ftype == "u" {
		toReturn.Ftype = "-"
		perm = "-" + perm[1:]
		toReturn.Permission = perm
	} else {
		toReturn.Ftype = ftype
		toReturn.Permission = perm
	}
	toReturn.OriginalName = file.Name()
	toReturn.Name = file.Name()
	if file.IsDir() { //si dossier
		path = append(path, path[0]+"/"+file.Name())
	} else if toReturn.Ftype == "l" { //si lien
		linkPath := path[0] + "/" + file.Name()
		templ := linkPath
		if linkPath[0] != '/' {
			linkPath = path[0] + "/" + linkPath
		}
		fileInfo, _ := os.Lstat(linkPath)
		if fileInfo.Mode()&os.ModeSymlink != 0 { // recherche si lien
			temp, _ := os.Readlink(linkPath)
			linkPath = temp
		} else {
			linkPath = templ
		}
		toReturn.Linkname = linkPath
	}
	if sys := file.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			//recherch major, minor, nlink, nom groupe et nom user
			toReturn.Major, toReturn.Minor = Annexe.DeviceNumber(stat.Rdev)
			toReturn.Link = int(stat.Nlink)
			groupName, _ := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
			toReturn.User = groupName.Name
			userName, _ := user.LookupId(strconv.Itoa(int(stat.Uid)))
			toReturn.Group = userName.Username
		}
	}
	return toReturn, path
}

func AddDots(arg []string, folderPath string, listeFetD []Struct.FileItem, r bool) []Struct.FileItem {
	ftype := ""
	perm := ""
	if r {
		doubleDots, _ := os.Stat(arg[0] + "/..")
		perm, ftype = Annexe.CheckPermission(doubleDots.Mode(), arg[0]+"/..")
		tempListeFetD, _ := StructStorage(doubleDots, perm, ftype, arg)
		tempListeFetD.FolderPath = folderPath // chemin du dossier contenant le fichier
		listeFetD = append(listeFetD, tempListeFetD)
	}
	dot, _ := os.Stat(arg[0] + "/.")
	perm, ftype = Annexe.CheckPermission(dot.Mode(), arg[0]+"/.")
	tempListeFetD, _ := StructStorage(dot, perm, ftype, arg)
	tempListeFetD.FolderPath = folderPath // chemin du dossier contenant le fichier
	listeFetD = append(listeFetD, tempListeFetD)
	if !r {
		doubleDots, _ := os.Stat(arg[0] + "/..")
		perm, ftype = Annexe.CheckPermission(doubleDots.Mode(), arg[0]+"/..")
		tempListeFetD, _ = StructStorage(doubleDots, perm, ftype, arg)
		tempListeFetD.FolderPath = folderPath // chemin du dossier contenant le fichier
		listeFetD = append(listeFetD, tempListeFetD)
	}
	return listeFetD
}
