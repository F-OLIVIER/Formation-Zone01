package lem_in

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fonction qui récupére les données du fichier d'entrée
func ParseFile(filename string) (int, []Room, []Link, []byte) {

	var data int
	var rooms []Room
	var links []Link

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Invalid file")
	}

	arr_lines := strings.Split(string(file), "\n")

	for i := 0; i < len(arr_lines); i++ {
		var room Room
		var attributes []string

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
		case arr_lines[i] == "##start":
			attributes = strings.Split(string(arr_lines[i+1]), " ")
			room.Name = attributes[0]
			room.X, _ = strconv.Atoi(attributes[1])
			room.Y, _ = strconv.Atoi(attributes[2])
			room.r_type = "start"
			i++
		case arr_lines[i] == "##end":
			attributes = strings.Split(string(arr_lines[i+1]), " ")
			room.Name = attributes[0]
			room.X, _ = strconv.Atoi(attributes[1])
			room.Y, _ = strconv.Atoi(attributes[2])
			room.r_type = "end"
			i++
		case strings.Contains(arr_lines[i], "-"): // if link
			attributes = strings.Split(string(arr_lines[i]), "-")
			links = append(links, Link{
				Current: attributes[0],
				Target:  attributes[1],
			})

		case arr_lines[i][0] == '#': // if comment
			continue

		default:
			attributes = strings.Split(string(arr_lines[i]), " ")
			room.Name = attributes[0]
		}

		if room.Name != "" && room.Name[0] != 'L' && room.Name[0] != '#' {
			room.X, _ = strconv.Atoi(attributes[1])
			room.Y, _ = strconv.Atoi(attributes[2])
			if room.r_type == "" {
				room.r_type = "intermediary"
			}
			rooms = append(rooms, room)
		}
	}
	return data, rooms, links, file
}
