package lem_in

import "fmt"

// Print la matrice des liens entre les rooms
func Print_Matrix(matrixed []Matrix) {
	tmp := matrixed[0].Name_X
	var order []string
	tmp_order := ""
	for _, v := range matrixed {
		if tmp_order != v.Name_X {
			order = append(order, v.Name_X)
			tmp_order = v.Name_X
		}
	}
	fmt.Printf("Taille de la matrice : ")
	fmt.Print(len(matrixed))
	fmt.Print(" cases\nPrint de la Matrice :\n  ")
	for _, v := range order {
		fmt.Printf(v + " ")
	}
	fmt.Println()
	count := 0
	fmt.Printf(order[count] + " ")
	for _, v := range matrixed {
		if tmp != v.Name_X {
			fmt.Println()
			count++
			fmt.Printf(order[count] + " ")
			tmp = v.Name_X
		}
		if !v.Linked {
			fmt.Printf(". ")
		} else {
			fmt.Printf("t ")
		}
	}
	fmt.Printf("\n\n")
}

// Print les routes
func PrintRoad(road_final []int, full_list_Path [][]string) {
	fmt.Println("\nRoutes selectionner :")
	for i, v := range road_final {
		fmt.Printf("Route n°%v : ", i+1)
		for j, value := range full_list_Path[v] {
			if j == 0 {
				fmt.Printf("%v", value)
			} else {
				fmt.Printf(" -> %v", value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Print une fourmis en ascii-art
func PrintFourmis() {
	res := "               ,\n"
	res += "      _,-'\\   /|   .    .    /`.\n"
	res += "  _,-'     \\_/_|_  |\\   |`. /   `._,--===--.__\n"
	res += " ^       _/\"/  \" \\ : \\__|_ /.   ,'    :.  :. .`-._\n"
	res += "        // ^   /7 t'\"\"    \"`-._/ ,'\\   :   :  :  .`.\n"
	res += "        Y      L/ )\\         ]],'   \\  :   :  :   : `.\n"
	res += "        |        /  `.n_n_n,','\\_    \\ ;   ;  ;   ;  _>\n"
	res += "        |__    ,'     |  \\`-'    `-.__\\_______.==---'\n"
	res += "       //  `\"\"\\\\      |   \\            \\\n"
	res += "       \\|     |/      /    \\            \\\n"
	res += "                     /     |             `.\n"
	res += "                    /      |               ^\n"
	res += "                   ^       |\n"
	res += "                           ^"
	fmt.Println(res)
}
