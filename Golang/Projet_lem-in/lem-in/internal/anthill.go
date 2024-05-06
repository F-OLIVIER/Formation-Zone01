package lem_in

import (
	"fmt"
	"os"
)

// ------------------------------------------------------
// ------------------ Matrice des liens -----------------
// ------------------------------------------------------

func InitMatrix(rooms []Room, links []Link) []Matrix {
	var matrixed []Matrix
	for i := 0; i < len(rooms); i++ {
		var tmp_matrix Matrix
		for j := 0; j < len(rooms); j++ {
			r_i, r_j := rooms[i].Name, rooms[j].Name
			tmp_matrix.Name_X, tmp_matrix.Name_Y = r_i, r_j

			for _, link := range links {
				if r_j == link.Current {
					if r_i == link.Target {
						tmp_matrix.Linked = true
						break
					}
				}
			}
			matrixed = append(matrixed, tmp_matrix)
			tmp_matrix.Linked = false
		}
	}
	return matrixed
}

// ------------------------------------------------------
// -------------- Fonction de backatraking --------------
// ------------------------------------------------------

func SearchPath(matrixed []Matrix, rooms []Room, name_salle_end string, list_Path []string, full_list_Path [][]string) ([]string, [][]string) {
	// Récupération des salles possible en fonction de la salle_current
	var salle_intermediare []string
	last_room := list_Path[len(list_Path)-1]
	for i := 0; i < len(matrixed); i++ {
		// sens normal
		if matrixed[i].Name_Y == last_room && matrixed[i].Linked {
			salle_intermediare = append(salle_intermediare, matrixed[i].Name_X)
		}
		// sens inversé
		if matrixed[i].Name_X == last_room && matrixed[i].Linked {
			salle_intermediare = append(salle_intermediare, matrixed[i].Name_Y)
		}
	}

	// check des chemins pour chaque salle intermédaire
	var tmp_list_Path []string
	tmp_list_Path = append(tmp_list_Path, list_Path...)
	// tmp_list_Path := list_Path
	for _, v := range salle_intermediare {
		if !checkRomm(tmp_list_Path, v) {
			// test de la salle courante (v)
			tmp_list_Path = append(tmp_list_Path, v)
			tmp_list_Path, full_list_Path = SearchPath(matrixed, rooms, name_salle_end, tmp_list_Path, full_list_Path)
			if v == name_salle_end { // gestion si salle end dans la liste, l'ajouter et retourner le résultats
				if !checkRomm(list_Path, name_salle_end) {
					full_list_Path = append(full_list_Path, tmp_list_Path)
					return nil, full_list_Path
				}
			}

			// reinitialisation de tmp_list_Path pour testé la salle suivante
			tmp_list_Path = list_Path
		}
	}
	return list_Path, full_list_Path
}

// ------------------------------------------------------
// ------------------ Fonction de check -----------------
// ------------------------------------------------------

// Fonction qui check la validation des salles "start" et "end"
func CheckData(matrixed []Matrix, name_salle_start, name_salle_end string) {
	var start bool
	var end bool
	for i := 0; i < len(matrixed); i++ {
		if matrixed[i].Name_X == name_salle_start || matrixed[i].Name_Y == name_salle_start {
			if matrixed[i].Linked {
				start = true
			}
		}
		if matrixed[i].Name_X == name_salle_end || matrixed[i].Name_Y == name_salle_end {
			if matrixed[i].Linked {
				end = true
			}
		}
	}
	if !start || !end {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("Absence of link at the \"start\" and/or \"end\" rooms")
		fmt.Println("Absence de lien au niveau des salles \"start\" et/ou \"end\"")
		os.Exit(0)
	}

}

// Fonction qui check chaque route et enregistre dans un map les route qui n'ont pas de point commum avec la route key
func CheckCommonPath(full_list_Path [][]string) map[int][]int {
	var select_Path [][]string
	var first_validate []int
	list_rooms := make(map[int][]int)
	select_Path = append(select_Path, full_list_Path[0])
	for i := 0; i < len(full_list_Path); i++ {
		current_path := full_list_Path[i]          // defini current_path "i"
		for j := 0; j < len(full_list_Path); j++ { // defini path_a_test "i+1"
			if i != j {
				path_a_test := full_list_Path[j]
				compareSS := compareSousSlice(current_path, path_a_test)
				if compareSS {
					first_validate = append(first_validate, j)
				}
			}
		}
		list_rooms[i] = first_validate
		first_validate = nil
	}
	return list_rooms
}

// Fonction qui vérification si la salle est deja dans la liste (évite de tournée en boucle), retourne true si deja presente
func checkRomm(list_Path []string, name_salle string) bool {
	if len(list_Path) == 0 {
		return false
	}
	for _, v := range list_Path {
		if v == name_salle {
			return true
		}
	}
	return false
}

// ------------------------------------------------------
// --------------- Fonction de sélection ----------------
// ------------------------------------------------------

// Fonction qui selectionne les routes à utiliser
func SelectAndSortRoad(list_Rooms map[int][]int) []int {
	var road_final [][]int
	var roadSelect []int
	for route_principal := 0; route_principal < len(list_Rooms); route_principal++ { // Parcourir la map (i est la key)
		current_road := list_Rooms[route_principal]
		// fmt.Println("current_road : ", current_road)
		if len(current_road) == 0 { // si aucun point commum avec les autres route (pas de vérification à effectuer)
			roadSelect = append(roadSelect, route_principal)
			road_final = append(road_final, roadSelect)
			roadSelect = nil
			continue
		} else if len(current_road) == 1 { // si une seule route sans point commum (pas de vérification à effectuer)
			roadSelect = append(roadSelect, route_principal)
			roadSelect = append(roadSelect, current_road[0])
			road_final = append(road_final, roadSelect)
			roadSelect = nil
			continue
		} else { // vérification à effectuer
			for range current_road { // parcourir les routes qui n'ont pas de point commum avec la route_principal
				roadSelect = append(roadSelect, route_principal)
				for _, element_current_road2 := range current_road {
					if compare_2_road(roadSelect, list_Rooms[element_current_road2]) {
						roadSelect = append(roadSelect, element_current_road2)

					}
				}
				road_final = append(road_final, roadSelect)
				roadSelect = nil
			}
		}
	}

	// Trie des routes avant retour pour renvoyer le maximum de routes possible
	road_final = sortFinalRoad(road_final)
	return road_final[0]
}

// ------------------------------------------------------
// -------------- Fonction de comparaison ---------------
// ------------------------------------------------------

// Fonction qui compare 2 liste et retourne true list_element_current_road a tout les éléments de roadSelect
func compare_2_road(roadSelect []int, list_element_current_road []int) bool {
	count := 0
	for _, value_roadSelect := range roadSelect {
		for _, value_current_road := range list_element_current_road {
			if value_current_road == value_roadSelect {
				count++
			}
		}
		if count == len(roadSelect) {
			return true
		}
	}
	return false
}

// Compare 2 slices et retourne true si aucune salle intermédaire en commum
func compareSousSlice(current_path []string, path_a_test []string) bool {
	if len(current_path) == 1 || len(path_a_test) == 1 {
		return true
	}
	for i := 1; i < len(current_path)-1; i++ {
		for j := 1; j < len(path_a_test)-1; j++ {
			if current_path[i] == path_a_test[j] {
				return false
			}
		}
	}
	return true
}

// ------------------------------------------------------
// ------------------ Fonction de trie ------------------
// ------------------------------------------------------

// Fonction qui trie les routes du plus petit vers le plus grand
func sortFinalRoad(list_road [][]int) [][]int {
	for i := 0; i < len(list_road); i++ {
		for j := 0; j < len(list_road)-1; j++ {
			if len(list_road[j]) < len(list_road[j+1]) {
				list_road[j], list_road[j+1] = list_road[j+1], list_road[j]
			}
		}
	}
	return list_road
}

// Fonction qui trie les routes du plus petit vers le plus grand
func Sort_full_list_Path(full_list_Path [][]string) [][]string {
	for i := 0; i < len(full_list_Path)-1; i++ {
		for j := 0; j < len(full_list_Path)-1; j++ {
			if len(full_list_Path[j]) > len(full_list_Path[j+1]) {
				full_list_Path[j], full_list_Path[j+1] = full_list_Path[j+1], full_list_Path[j]
			}
		}
	}
	return full_list_Path
}

// Fonction qui trie les routes du plus grand vers le plus petit
func SortFinalSubRoad(road_final []int, full_list_Path [][]string) []int {
	for i := 0; i < len(road_final); i++ {
		for j := 0; j < len(road_final)-1; j++ {
			if len(full_list_Path[road_final[j]]) > len(full_list_Path[road_final[j+1]]) {
				road_final[j], road_final[j+1] = road_final[j+1], road_final[j]
			}
		}
	}
	return road_final
}

// place la salle "start" au debut et la salle "end" à la fin de la liste
func SortRomms(rooms []Room) []Room {
	// for i := 0; i < len(rooms); i++ {
	for j := 0; j < len(rooms)-1; j++ {
		if rooms[j].r_type == "start" {
			rooms[j], rooms[0] = rooms[0], rooms[j]
		}
		if rooms[j].r_type == "end" {
			rooms[j], rooms[len(rooms)-1] = rooms[len(rooms)-1], rooms[j]
		}
	}
	// }
	return rooms
}

// ------------------------------------------------------
// ------------------- Autres fonction ------------------
// ------------------------------------------------------

// Fonction d'initialisation pour les salles "start" et "end"
func Start_end(rooms []Room, list_Path []string) ([]string, string, string) {
	var name_salle_end string
	var name_salle_start string
	for _, v := range rooms { // Initialisation de list_Path avec le name de la salle start
		if v.r_type == "start" {
			name_salle_start = v.Name
			list_Path = append(list_Path, name_salle_start)
		}
		if v.r_type == "end" { // Récupération name salle end
			name_salle_end = v.Name
		}
	}
	return list_Path, name_salle_start, name_salle_end
}
