package lem_in

import (
	"fmt"
	"strconv"
)

func SendAnts(nb_ants int, road_final []int, full_list_Path [][]string) {
	var list_road [][]string
	var max_lenght int
	min_lenght := len(full_list_Path[0])
	for _, v := range road_final {
		list_road = append(list_road, full_list_Path[v])
		if max_lenght < len(full_list_Path[v]) {
			max_lenght = len(full_list_Path[v])
		}
	}

	// Définition de la route a emprunter pour les fourmis
	nb_send := 1
	var list_ants []Ant
	for i := 1; i < max_lenght; i++ { // parcour chaque route
		tmp_list_road := list_road
		for nb_ants >= nb_send {
			if nb_ants-nb_send >= max_lenght-min_lenght { // si le nombre de fourmis est supérieur a la taille de la route
				for j := 0; j < len(tmp_list_road); j++ { // récupére le nom de la salle
					current_road := tmp_list_road[j]
					if i < len(current_road) { // Condition pour éviter le depassement en taille d'une route
						var ant Ant
						ant.Name = "L" + strconv.Itoa(nb_send) // nom de la fourmis
						ant.Room = current_road[i]             // Salle de depart de la fourmis
						ant.Road = j                           // chemin de la fourmis
						ant.Road_lenght = len(current_road)    // taille de la route
						ant.Position = 0                       // emplacement sur la route
						list_ants = append(list_ants, ant)
					}
					if nb_send <= nb_ants { //Condition d'arrete d'envoi de fourmis
						nb_send++
					}
				}
			} else { // si le nombre de fourmis est inférieur à la taille de la route, on utilise que le chemin n°0
				current_road := tmp_list_road[0]
				if i < len(current_road) { // Condition pour éviter le depassement en taille d'une route
					var ant Ant
					// fmt.Println("HERE : L" + strconv.Itoa(nb_send))
					ant.Name = "L" + strconv.Itoa(nb_send) // nom de la fourmis
					ant.Room = current_road[i]             // Salle de depart de la fourmis
					ant.Road = 0                           // chemin de la fourmis
					ant.Road_lenght = len(current_road)    // taille de la route
					ant.Position = 0                       // emplacement sur la route
					list_ants = append(list_ants, ant)
				}
				if nb_send <= nb_ants { //Condition d'arrete d'envoi de fourmis
					nb_send++
				}
				// break
			}
		}
	}

	nb_tour := 1
	// Boucle de print et d'avancement des fourmis
	for {
		fmt.Printf("Move %v : ", nb_tour)
		var room_visited []string
		var road_visited []int
		for i := 0; i < len(list_ants); i++ {
			var current Ant
			current = list_ants[i]
			if !checkVisited(current.Room, room_visited) { // si salle non visité
				if list_ants[i].Position+1 < list_ants[i].Road_lenght-1 { // déplacement de la fourmis
					fmt.Printf("%v-%v ", current.Name, current.Room)
					room_visited = append(room_visited, current.Room)
					current.Position += 1
					current.Room = list_road[current.Road][current.Position+1]
				} else {
					if current.Arret == false {
						// vérification si le chemin fais start-end
						if !checkroadVisited(current.Road, road_visited) {
							fmt.Printf("%v-%v ", current.Name, current.Room)
							current.Arret = true
							road_visited = append(road_visited, current.Road)
						}
					}
				}
				list_ants[i] = current
			}
		}
		fmt.Println()
		if len(room_visited) == 0 { // condition d'arret
			break
		} // re-initialisation des valeur avant la boucle suivante
		room_visited = nil
		road_visited = nil
		nb_tour++
	}
	// fmt.Printf("\nNombre de tours pour faire passer toutes les fourmis : %v\n", nb_tour)
}

// Fonction qui renvoi true si la salle à été visité
func checkVisited(room string, room_visited []string) bool {
	for _, current_room := range room_visited {
		if room == current_room {
			return true
		}
	}
	return false
}

// Fonction qui renvoi true si la route à déja été visité
func checkroadVisited(road int, road_visited []int) bool {
	for _, current_road := range road_visited {
		if road == current_road {
			return true
		}
	}
	return false
}
