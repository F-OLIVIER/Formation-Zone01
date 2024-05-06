package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	Annexe "my-ls-1/package/annexe"
	Check "my-ls-1/package/check"
	LS "my-ls-1/package/ls"
	Sort "my-ls-1/package/sort"
	Struct "my-ls-1/package/struct"
)

// exemple de command : go run . -l package -a audit.sh
func main() {
	listToPrint := [][]Struct.FileItem{}
	folderList := []string{}
	args := os.Args[1:]
	listeOption := Check.CheckEntrer(args)
	if len(args) == 0 { // si args vide
		// fonction Ls(l bool, r bool, R bool, a bool, t bool, folderList[]string)
		folderList = append(folderList, ".")
		tempOption := []Struct.Option{Struct.Option{false, false, false, false, false, []string{"."}}}
		for _, v := range folderList {
			listToPrint = append(listToPrint, rechercheFile(v, listToPrint, tempOption, 0)...)
		}
		for _, v := range listToPrint {
			Annexe.Printlist(v, false, false, false, false, false)
		}
		// listToPrint = append(listToPrint, Annexe.Ls("/dev", true, false, true, false, true))
	} else {
		for i := 0; i < len(listeOption); i++ {
			if len(listeOption[i].Argument) == 0 { // si les arguments dans args vide
				folderList = append(folderList, ".")
				for _, v := range folderList {
					listToPrint = append(listToPrint, rechercheFile(v, listToPrint, listeOption, i)...)
				}
				aRemplacer := folderList[0]
				if listeOption[i].OptionGrandR {
					listToPrint = Sort.InsertionSort(listToPrint, listeOption[i].OptionPetitR, listeOption[i].OptionPetitT)
				}
				listToPrint = calculTotal(listToPrint, listeOption[i].OptionPetitA)
				for _, v := range listToPrint {
					if listeOption[i].OptionGrandR {
						pathAprint := strings.Replace(v[0].FolderPath, aRemplacer, ".", 1)
						fmt.Printf("\n%v:\n", pathAprint)
					}
					if listeOption[i].OptionPetitL {
						fmt.Printf("total %v\n", v[0].Total)
					}
					Annexe.Printlist(v, listeOption[i].OptionPetitL, listeOption[i].OptionPetitR, listeOption[i].OptionGrandR, listeOption[i].OptionPetitA, listeOption[i].OptionPetitT)
				}
			} else {
				for j := 0; j < len(listeOption[i].Argument); j++ {
					folderList = append(folderList, listeOption[i].Argument[j])
				}
				for _, v := range folderList {
					listToPrint = rechercheFile(v, listToPrint, listeOption, i)
				}
				if listeOption[i].OptionGrandR {
					listToPrint = Sort.InsertionSort(listToPrint, listeOption[i].OptionPetitR, listeOption[i].OptionPetitT)
				}
				listToPrint = calculTotal(listToPrint, listeOption[i].OptionPetitA)
				for itt, v := range listToPrint {
					// Gestion affichage du nom de dossier si option "R"
					if (listeOption[i].OptionGrandR || listeOption[i].OptionPetitL) && (len(folderList) > 1 || len(listToPrint) > 1) || listeOption[i].OptionGrandR {
						path := v[0].FolderPath
						if itt > 0 {
							fmt.Printf("\n%v:\n", path)
						} else {
							fmt.Printf("\n%v:\n", path)
						}
					}
					if listeOption[i].OptionPetitL {
						fmt.Printf("total %v\n", v[0].Total)
					}
					Annexe.Printlist(v, listeOption[i].OptionPetitL, listeOption[i].OptionPetitR, listeOption[i].OptionGrandR, listeOption[i].OptionPetitA, listeOption[i].OptionPetitT)
				}
			}
		}
	}
}
func rechercheFile(folder string, listToPrint [][]Struct.FileItem, listeOption []Struct.Option, i int) [][]Struct.FileItem {
	folderList := []string{folder}
	count := 0
	for len(folderList) > 0 {
		var tempListToPrint []Struct.FileItem
		tempListToPrint, folderList = LS.Ls(listeOption[i].OptionPetitL, listeOption[i].OptionPetitR, listeOption[i].OptionGrandR, listeOption[i].OptionPetitA, listeOption[i].OptionPetitT, folderList, count)
		// fmt.Println("-------------test-------------- : ", folderList)
		if tempListToPrint != nil {
			listToPrint = append(listToPrint, tempListToPrint)
		}
		// for _, v := range listToPrint[0] {
		// 	fmt.Printf("v[0].Name : %v\n", v.Name)
		// }
		count++
	}
	return listToPrint
}
func calculTotal(listToPrint [][]Struct.FileItem, a bool) [][]Struct.FileItem {
	for i := len(listToPrint) - 1; i >= 0; i-- {
		var total int64
		for _, k := range listToPrint[i] {
			if !a && k.OriginalName[0] == '.' { // Retirer les fichiers caché, . et ..
			} else {
				if k.Ftype == "l" { // Retirer les liens
				} else {
					fileInfo, _ := os.Stat(k.FolderPath + "/" + k.OriginalName)
					sysInfo, _ := fileInfo.Sys().(*syscall.Stat_t)
					total = sysInfo.Blocks // Nombre de blocs utilisés
					listToPrint[i][0].Total += int(total / 2)
				}
			}
		}
		total = 0
	}
	return listToPrint
}
