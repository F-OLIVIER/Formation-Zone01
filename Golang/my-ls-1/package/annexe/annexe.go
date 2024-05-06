package Annexe

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"syscall"
	"time"

	Struct "my-ls-1/package/struct"
)

const _LISTXATTR_BUFFER = 1024

func CheckPermission(perm fs.FileMode, arg string) (permissionClean string, fileType string) {
	permString := perm.String()
	// in order:
	// if 11 char remove 1 left
	// if start c and xattr size : 24 and string(xattrBuffer[:xattrSize]) == system.posix_acl_access
	// then +
	if len(permString) == 11 {
		fileType = string(permString[1])
		permissionClean = string(permString[1:])
		xattrBuffer := make([]byte, _LISTXATTR_BUFFER)
		xattrSize, _ := syscall.Listxattr(arg, xattrBuffer)
		if xattrSize > 0 {
			permissionClean += "+"
		}
	} else {
		fileType = string(permString[0])
		permissionClean = string(permString)
	}
	return permissionClean, fileType
}

func DeviceNumber(dev uint64) (uint32, uint32) {
	major := uint32((dev & 0x00000000000fff00) >> 8)
	major |= uint32((dev & 0xfffff00000000000) >> 32)
	minor := uint32((dev & 0x00000000000000ff))
	minor |= uint32((dev & 0x00000ffffff00000) >> 12)
	return major, minor
}

func tradmois(numMois int) (mois string) {
	if numMois == 1 {
		mois = "janv."
	} else if numMois == 2 {
		mois = "févr."
	} else if numMois == 3 {
		mois = "mars"
	} else if numMois == 4 {
		mois = "avril"
	} else if numMois == 5 {
		mois = "mai"
	} else if numMois == 6 {
		mois = "juin"
	} else if numMois == 7 {
		mois = "juil."
	} else if numMois == 8 {
		mois = "août"
	} else if numMois == 9 {
		mois = "sept."
	} else if numMois == 10 {
		mois = "oct."
	} else if numMois == 11 {
		mois = "nov."
	} else if numMois == 12 {
		mois = "déc."
	} else if numMois == 13 {
		mois = ""
	}
	return mois
}

// Formate le name pour ignorer les caractéres spéciaux des noms pour la comparaison, fait partie de sortSlice
func TraitementName(name string) string {
	tmp := ""
	for _, i := range name {
		if (i >= '0' && i <= '9') || (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			tmp += string(i)
		}
	}
	if len(tmp) != 0 { // gestion des noms vide  apres retrait des caractéres spéciaux
		name = tmp
	}
	// fmt.Printf("name : %v \t\t\ttmp : %v\n", name, tmp)
	return strings.ToLower(name)
}

func CherchePoint(s string) bool {
	if s == "./" {
		s = "."
	} else if s == "../" {
		s = ".."
	}
	for i := len(s) - 1; i > 0; i-- {
		if i < len(s)-2 {
			if s[i] == '/' && s[i+1] == '.' {
				return true
			}
		}
	}
	return false
}

// Check si le fileName contient des caractéres spéciaux et retourne le fileName avec les quotes adapté (simple ou double)
func testContaintForQuote(fileName string) string {
	for _, car := range fileName {
		j := string(car)
		if j == " " ||
			j == "[" ||
			j == "]" ||
			j == "^" ||
			j == "(" ||
			j == ")" ||
			j == "{" ||
			j == "}" ||
			j == "|" ||
			j == "`" ||
			j == "#" ||
			j == "&" ||
			j == "~" ||
			j == "\\" ||
			j == "\"" {
			fileName = "'" + fileName + "'"
			return fileName
		} else if j == "'" {
			fileName = "\"" + fileName + "\""
			return fileName
		}
	}
	return fileName
}

func Printlist(fileList []Struct.FileItem, l bool, r bool, R bool, a bool, t bool) {
	for _, file := range fileList {
		if !a && file.OriginalName[0] == '.' {
			// si pas l'option "a" et fichier caché (commencant par "."), on fais rien
		} else {
			if l { // si option "l"
				var date string
				theTime := time.Date(file.Lastmod.Year(), file.Lastmod.Month(), file.Lastmod.Day(), file.Lastmod.Hour(), file.Lastmod.Minute(), file.Lastmod.Second(), 100, time.Local)
				if file.Lastmod.Year() == time.Now().Year() && // si années identique et moins de 7 mois de différence on affiche l'heure sinon la date
					int(file.Lastmod.Month()) > int(time.Now().Month())-7 {
					date = theTime.Format("02 15:04")
				} else {
					date = theTime.Format("02  2006")
				}
				tmpFileName := testContaintForQuote(file.Name)
				tmpLinkName := testContaintForQuote(file.Linkname)
				if file.Major > 0 || file.Minor > 0 { // Print pour les fichiers spéciaux
					if len(tmpLinkName) > 0 {
						fmt.Printf("%-11v %3v %-6v %-8v %3v, %5v %-5v %v %v -> %v\n", file.Permission, file.Link, file.Group, file.User, strconv.Itoa(int(file.Major)), strconv.Itoa(int(file.Minor)), tradmois(int(theTime.Month())), date, tmpFileName, tmpLinkName)
					} else {
						fmt.Printf("%-11v %3v %-6v %-8v %3v, %5v %-5v %v %v\n", file.Permission, file.Link, file.Group, file.User, strconv.Itoa(int(file.Major)), strconv.Itoa(int(file.Minor)), tradmois(int(theTime.Month())), date, tmpFileName)
					}
				} else { // Print option -l
					if len(tmpLinkName) > 0 {
						fmt.Printf("%-11v %3v %-6v %-8v %10v %-5v %v %v -> %v\n", file.Permission, file.Link, file.Group, file.User, file.Size, tradmois(int(theTime.Month())), date, tmpFileName, tmpLinkName)
					} else {
						fmt.Printf("%-11v %3v %-6v %-8v %10v %-5v %v %v\n", file.Permission, file.Link, file.Group, file.User, file.Size, tradmois(int(theTime.Month())), date, tmpFileName)
					}
				}
			} else {
				if file.Name[0] != '.' { // Print par defaut (sans option) et -a et -r et -R et -t
					// Enregistrement dans listeAprint pour print l'ensemble d'un coup aprés modification
					file.Name = testContaintForQuote(file.Name)
					fmt.Printf("%v\n", file.Name)
				} else if a && file.Name[0] == '.' {
					// Print les fichiers "." caché (option -a)
					file.Name = testContaintForQuote(file.Name)
					file.Linkname = testContaintForQuote(file.Linkname)
					fmt.Printf("%v\n", file.Name)
				}
			}
		}
	}
}
