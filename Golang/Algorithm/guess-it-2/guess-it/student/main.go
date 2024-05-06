package main

import (
	"fmt"
)

// command : node server.js
func main() {
	var numY float64
	var dataX []float64
	var dataY []float64
	numX := 0
	for a := 1; a <= 12500; a++ {
		//Récupération de l'entrée
		fmt.Scan(&numY)
		dataX = append(dataX, float64(numX))
		numX++
		dataY = append(dataY, numY)

		// Prediction du résultat
		a, b := LinearRegressionLine(dataX, dataY)
		y := a*float64(numX) + b

		// définition de la plage
		valeurMin := y - 29
		valeurMax := y + 29
		// valeurMin := y + r
		// valeurMax := y - r
		fmt.Println(valeurMin, valeurMax)
	}
}

// Fonction de calcul de la ligne de régression linéaire  (rendu a et b parmis : y = ax + b)
func LinearRegressionLine(datax []float64, datay []float64) (penteA float64, interceptionB float64) {
	// moyenne des data
	moyenneDataX := Average(datax)
	moyenneDataY := Average(datay)
	var quadratureDatax float64
	var sommeXY float64
	for i := 0; i < len(datax); i++ {
		x := datax[i]
		y := datay[i]
		// Quadrature de la différence et addition des résultats pour X
		quadratureDatax = quadratureDatax + ((moyenneDataX - x) * (moyenneDataX - x))
		// Multiplication entre Avg(X)-X et Avg(Y)-Y et somme des résultats
		sommeXY = sommeXY + ((moyenneDataX - x) * (moyenneDataY - y))
	}
	penteA = sommeXY / quadratureDatax
	interceptionB = moyenneDataY - (penteA * moyenneDataX)
	return penteA, interceptionB
}

// Moyenne
func Average(table []float64) (res float64) {
	var temp float64
	for a := 0; a < len(table); a++ {
		temp = temp + table[a]
	}
	res = temp / float64(len(table))
	return res
}
