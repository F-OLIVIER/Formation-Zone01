package main

import (
	"fmt"
	"sort"
)

// command : node server.js

func main() {
	var num float64
	var data []float64

	for a := 1; a <= 12500; a++ {
		//Récupération de l'entrée
		fmt.Scan(&num)

		data = append(data, num)
		// trie du tableau au fur et a mesure pour le calcul de la medianne
		sort.Float64s(data)
		med := Median(data)

		// définition de la plage
		valeurMin := med - 29
		valeurMax := med + 29

		fmt.Println(valeurMin, valeurMax)
	}
}

// Médiane
func Median(table []float64) int {
	taille := len(table)
	if taille%2 == 0 { // taille pair
		return int((table[taille/2] + table[taille/2-1]) / 2)
	} else { // taille impair
		return int(table[taille/2])
	}
}
