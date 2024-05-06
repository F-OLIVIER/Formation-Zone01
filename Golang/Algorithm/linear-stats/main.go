package main

import (
	"fmt"
	"os"

	C "linear-stats/packages/calcul-stats"
	F "linear-stats/packages/file"
)

// command go run . data.txt

func main() {
	// debut := time.Now()

	if len(os.Args[1:]) != 1 {
		fmt.Println("ERROR : probléme dans le nombre d'argument de la commande")
		fmt.Println("Exemple de commande : go run your-program.go data.txt")
		os.Exit(1)
	}

	// Ouverture du fichier et récupération des data
	datax, datay := F.OpenFile(os.Args[1])
	fmt.Println("datax : ", datax)

	// Calcul de la ligne de régression linéaire
	a, b := C.LinearRegressionLine(datax, datay)
	fmt.Printf("Linear Regression Line: y = ")
	fmt.Printf("%.6f", a)
	fmt.Printf("x + ")
	fmt.Printf("%.6f", b)
	fmt.Println()

	// Calcul du coefficient de corrélation de Pearson
	result := C.PearsonCorrelationCoefficent(datax, datay)
	fmt.Printf("Pearson Correlation Coefficient: ")
	fmt.Printf("%.10f", result)
	fmt.Println()

	// Affichage du temps de calcul
	// fin := time.Now()
	// fmt.Println()
	// fmt.Println(fin.Sub(debut))
}
