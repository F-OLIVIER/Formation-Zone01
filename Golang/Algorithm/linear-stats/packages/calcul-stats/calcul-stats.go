package calculstats

import (
	"math"
)

// Site d'aide pour le calcul :
// https://www.voxco.com/fr/blog/comment-calculer-la-regression-lineaire/#:~:text=Calcul%20de%20la%20régression%20linéaire,équation%20qui%20correspond%20aux%20données.&text=L%27équation%20se%20présente%20sous,comme%20la%20formule%20de%20pente.

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

// https://en.wikipedia.org/wiki/Pearson_correlation_coefficient

// Fonction de calcul du coefficient de corrélation de Pearson (ou coefficient r)
func PearsonCorrelationCoefficent(datax []float64, datay []float64) (result float64) {
	// moyenne des data
	moyenneDatax := Average(datax)
	moyenneDatay := Average(datay)

	// calcul des sommes
	var sommeXMoyenneCarre float64
	var sommeYMoyenneCarre float64
	var sommeXMoyenneMoinsYMoyenne float64
	for i := 0; i < len(datax); i++ {
		sommeXMoyenneMoinsYMoyenne = sommeXMoyenneMoinsYMoyenne + ((datax[i] - moyenneDatax) * (datay[i] - moyenneDatay))
		sommeXMoyenneCarre = sommeXMoyenneCarre + ((datax[i] - moyenneDatax) * (datax[i] - moyenneDatax))
		sommeYMoyenneCarre = sommeYMoyenneCarre + ((datay[i] - moyenneDatay) * (datay[i] - moyenneDatay))
	}

	result = sommeXMoyenneMoinsYMoyenne / (math.Sqrt(sommeXMoyenneCarre) * math.Sqrt(sommeYMoyenneCarre))
	return result
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
