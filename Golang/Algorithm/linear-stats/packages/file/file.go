package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Fonction d'ouverture du fichier et de renvoie des cordonnée x et y de chaque point
func OpenFile(nameFile string) (datax []float64, datay []float64) {
	// Ouverture du fichier
	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Printf("ERROR, probléme avec le nom du fichier : ")
		fmt.Println(err)
		os.Exit(1)
	}

	// Création d'un scanner
	sc := bufio.NewScanner(file)
	line := 0
	for sc.Scan() { // scan de la ligne courante
		// fmt.Println(sc.Text())
		datax = append(datax, float64(line))
		line++

		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("ERROR : probléme avec le contenu du fichier, ligne : ", line)
			os.Exit(1)
		}
		datay = append(datay, float64(num))
	}
	return datax, datay
}
