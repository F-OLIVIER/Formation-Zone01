package lem_in

import (
	"fmt"
)

func Travail(args []string) {
	if len(args) != 1 {
		fmt.Println("ERROR: invalid arguments")
		return
	}
	nb_ants, rooms, links, data_file := ParseFile(args[0])
	matrixed := InitMatrix(rooms, links)
	// Print de la Matrice généré
	// Print_Matrix(matrixed)
	var list_Path []string
	var full_list_Path [][]string
	var name_salle_end string
	var name_salle_start string
	rooms = SortRomms(rooms)                                                  // tri des salles mise de "start" au debut et "end" à la fin
	list_Path, name_salle_start, name_salle_end = Start_end(rooms, list_Path) // Ajout salle start a "list_Path" et récupération name salle "end"
	CheckData(matrixed, name_salle_start, name_salle_end)                     // Vérification des entrées (es-ce que les salles start et end sont connecté a une salle ?)
	fmt.Println(string(data_file))
	list_Path, full_list_Path = SearchPath(matrixed, rooms, name_salle_end, list_Path, full_list_Path) // récursive de récupération des chemins
	// fmt.Println("\nListe des routes possible :")
	// for i, v := range full_list_Path {
	// 	fmt.Println("chemin n°", i, ":", v)
	// }
	full_list_Path = Sort_full_list_Path(full_list_Path)
	list_Rooms := CheckCommonPath(full_list_Path)
	// fmt.Println("\nMap d'array des chemin sans point commum (ATTENTION, seul la key est vérifié) :")
	// for i := 0; i < len(list_Rooms); i++ {
	// 	fmt.Println("chemin n°", i, ":", list_Rooms[i])
	// }
	// fmt.Println()

	road_final := SelectAndSortRoad(list_Rooms)
	road_final = SortFinalSubRoad(road_final, full_list_Path)
	PrintRoad(road_final, full_list_Path)
	SendAnts(nb_ants, road_final, full_list_Path)

}
