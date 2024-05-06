package visualizer_lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fonction qui récupére le Stdin
func GetInput() (input []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// Fonction qui récupére les données du fichier d'entrée
func ParseFile(arr_lines []string) (int, []Room, []Link, []Ant) {

	var data int
	var rooms []Room
	var links []Link
	var ants []Ant
	var intermediare bool

	for i := 0; i < len(arr_lines); i++ {
		var room Room
		var attributes []string
		if len(arr_lines[i]) > 0 {
			// fmt.Println("test : ", arr_lines[i])
			if !intermediare {
				switch {
				case i == 0:
					var err_ant_nb error
					data, err_ant_nb = strconv.Atoi(arr_lines[i])
					if err_ant_nb != nil {
						fmt.Println("Ant total number is not an integer")
					}
					if data <= 0 {
						fmt.Println("ERROR: invalid data format")
						fmt.Println("Absence of ants")
						fmt.Println("Absence de foumis")
						os.Exit(0)
					}
				case strings.Contains(arr_lines[i], "Routes selectionner"):
					intermediare = true
					continue
				case arr_lines[i] == "##start":
					attributes = strings.Split(string(arr_lines[i+1]), " ")
					room.Name = attributes[0]
					room.X, _ = strconv.Atoi(attributes[1])
					room.Y, _ = strconv.Atoi(attributes[2])
					room.r_type = "start"
					rooms = append(rooms, room)
					i++

				case arr_lines[i] == "##end":
					attributes = strings.Split(string(arr_lines[i+1]), " ")
					room.Name = attributes[0]
					room.X, _ = strconv.Atoi(attributes[1])
					room.Y, _ = strconv.Atoi(attributes[2])
					room.r_type = "end"
					rooms = append(rooms, room)
					i++

				case arr_lines[i][0] == '#': // if comment
					continue

				case strings.Contains(arr_lines[i], "-"): // if link
					attributes = strings.Split(string(arr_lines[i]), "-")
					links = append(links, Link{
						Current: attributes[0],
						Target:  attributes[1],
					})

				default:
					attributes = strings.Split(string(arr_lines[i]), " ")
					room.Name = attributes[0]
					if room.Name != "" && room.Name[0] != 'L' && room.Name[0] != '#' {
						room.X, _ = strconv.Atoi(attributes[1])
						room.Y, _ = strconv.Atoi(attributes[2])
						if room.r_type == "" {
							room.r_type = "intermediary"
						}
						rooms = append(rooms, room)
					}
				}
			} else {
				var ant Ant
				switch {
				case strings.Contains(arr_lines[i], "Move"):
					attributes = strings.Split(string(arr_lines[i]), " ")
					numMove, err_atoi := strconv.Atoi(attributes[1])
					if err_atoi != nil {
						fmt.Println("err_atoi : ", err_atoi)
					}

					for _, v := range attributes {
						if strings.Contains(v, "-") {
							subAttributes := strings.Split(string(v), "-")
							ant.Move = numMove
							ant.Name = subAttributes[0]
							ant.Room = subAttributes[1]
							ants = append(ants, ant)
						}
					}
				default:
					continue
				}
			}
		}

	}

	return data, rooms, links, ants
}
