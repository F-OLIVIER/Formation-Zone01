package Sort

import (
	"strings"
	"time"

	Annexe "my-ls-1/package/annexe"
	Struct "my-ls-1/package/struct"
)

// Fonction de tri (option 1: ordre croissant - autre n°: ordre decroissant)
func SortSlice(listeFetD []Struct.FileItem, ordre int) []Struct.FileItem {
	if ordre == 0 {
		for i := 1; i < len(listeFetD); i++ {
			key := listeFetD[i]
			j := i - 1
			name2 := Annexe.TraitementName(key.OriginalName)
			for j >= 0 && strings.Compare(Annexe.TraitementName(listeFetD[j].OriginalName), name2) > 0 {
				listeFetD[j+1] = listeFetD[j]
				j = j - 1
			}
			listeFetD[j+1] = key
		}
	} else {
		for i := len(listeFetD) - 2; i >= 0; i-- {
			key := listeFetD[i]
			j := i + 1
			name2 := Annexe.TraitementName(key.OriginalName)
			for j < len(listeFetD) && strings.Compare(Annexe.TraitementName(listeFetD[j].OriginalName), name2) > 0 {
				listeFetD[j-1] = listeFetD[j]
				j = j + 1
			}
			listeFetD[j-1] = key
		}
	}
	return listeFetD
}

// Fonction de tri par date de modification (option -t)
func SortSliceDate(listeFetD []Struct.FileItem, r bool) []Struct.FileItem {
	for i := 0; i < len(listeFetD); i++ {
		for j := 0; j < len(listeFetD)-1; j++ {
			infoJ := listeFetD[j].Lastmod
			infoJJ := listeFetD[j+1].Lastmod
			if (infoJ.After(infoJJ) && r) || (infoJ.Before(infoJJ) && !r) {
				listeFetD[j], listeFetD[j+1] = listeFetD[j+1], listeFetD[j]
			} else if infoJ.Equal(infoJJ) {
				var temp []Struct.FileItem
				temp = append(temp, listeFetD[j])
				temp = append(temp, listeFetD[j+1])
				if r {
					temp = SortSlice(temp, 1)
				} else {
					temp = SortSlice(temp, 0)
				}
				listeFetD[j], listeFetD[j+1] = temp[0], temp[1]
			}
		}
	}
	return listeFetD
}

// Fonction qui trie "listToPrint"
func InsertionSort(listToPrint [][]Struct.FileItem, r bool, t bool) [][]Struct.FileItem {
	// retrait des . et / pour dans le FolderPath pour le trie
	for i, v := range listToPrint {
		for j, _ := range v {
			listToPrint[i][j].FolderPathT = strings.Replace(listToPrint[i][j].FolderPath, ".", "", -1)
			listToPrint[i][j].FolderPathT = strings.Replace(listToPrint[i][j].FolderPathT, "/", "", -1)
		}
	}

	// trie des structures "i"
	for k := 0; k < len(listToPrint); k++ { // boucle d'augmentation du nombre de trie
		for i := 0; i < len(listToPrint)-1; i++ {
			for ii := 0; ii < len(listToPrint)-1; ii++ {
				// définition de variable si "t" est demandé
				var infoJ time.Time
				var tmpFolderPathJ string
				var infoJJ time.Time
				var tmpFolderPathJJ string
				if t { // Récupération des information lastmod et FolderPathT du dossier "." si "t" est demandé
					for _, v := range listToPrint[ii] { // pour i
						if v.Name == "." {
							infoJ = v.Lastmod
							tmpFolderPathJ = v.FolderPathT
						}
					}
					for _, v := range listToPrint[ii+1] { // pour i+1
						if v.Name == "." {
							infoJJ = v.Lastmod
							tmpFolderPathJJ = v.FolderPathT
						}
					}
				}
				if r && t { // trie par date de modification inversé
					if !strings.Contains(tmpFolderPathJJ, tmpFolderPathJ) &&
						infoJJ.Before(infoJ) {
						listToPrint[ii], listToPrint[ii+1] = listToPrint[ii+1], listToPrint[ii]
					}
				} else if t { // trie par date de modification
					if !strings.Contains(tmpFolderPathJJ, tmpFolderPathJ) &&
						infoJ.Before(infoJJ) {
						listToPrint[ii], listToPrint[ii+1] = listToPrint[ii+1], listToPrint[ii]
					}
				} else if r { // trie par ordre alphabétique inversé
					if !strings.Contains(listToPrint[ii+1][0].FolderPathT, listToPrint[ii][0].FolderPathT) {
						for listToPrint[ii][0].FolderPathT < listToPrint[ii+1][0].FolderPathT {
							listToPrint[ii], listToPrint[ii+1] = listToPrint[ii+1], listToPrint[ii]
						}
					}
				}
			}
		}
	}

	// trie des sous structures "j" de "i"
	for k := 0; k < len(listToPrint); k++ {
		for i := 0; i < len(listToPrint); i++ { // boucle d'augmentation du nombre de trie
			for j := 0; j < len(listToPrint[i])-1; j++ {
				// définition de variable si "t" est demandé
				infoJ := listToPrint[i][j].Lastmod
				infoJJ := listToPrint[i][j+1].Lastmod

				if r && t { // trie par date de modification inversé
					if listToPrint[i][j].FolderPathT == listToPrint[i][j+1].FolderPathT {
						if infoJJ.Before(infoJ) {
							listToPrint[i][j], listToPrint[i][j+1] = listToPrint[i][j+1], listToPrint[i][j]
						}
					}
				} else if t { // trie par date de modification
					if listToPrint[i][j].FolderPathT == listToPrint[i][j+1].FolderPathT {
						if infoJ.Before(infoJJ) {
							listToPrint[i][j], listToPrint[i][j+1] = listToPrint[i][j+1], listToPrint[i][j]

						}
					}
				} else if r { // trie par ordre alphabétique inversé
					if listToPrint[i][j].FolderPathT == listToPrint[i][j+1].FolderPathT {
						for listToPrint[i][j].Name < listToPrint[i][j+1].Name {
							listToPrint[i][j], listToPrint[i][j+1] = listToPrint[i][j+1], listToPrint[i][j]
						}
					}
				}
			}
		}
	}

	return listToPrint
}
