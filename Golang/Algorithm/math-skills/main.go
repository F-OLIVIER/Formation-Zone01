package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

var verif = true

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 || len(arg) > 1 {
		fmt.Println("missing arguments")
		return
	}

	file := arg[0]
	// lecture du fichier d'entrée
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Tranformationde l'array de byte en array de int
	table := tranformInt(data)
	if verif == false {
		fmt.Println("Contenu du fichier incorrect")
		return
	}

	fmt.Printf("Average: ")
	print(Average(table))
	fmt.Printf("Median: ")
	print(Median(table))
	fmt.Printf("Variance: ")
	print(Variance(table))
	fmt.Printf("Standard Deviation: ")
	print(Standard_Deviation(table))
}

// Tranformationde l'array de byte en array de int
// int64 min : -9223372036854775808
// int64 max : 9223372036854775807
func tranformInt(data []byte) (table []float64) {
	var num float64
	negatif := false
	for a := 0; a < len(data); a++ {
		if data[a] >= '0' && data[a] <= '9' {
			num = num*10 + float64(data[a]-'0')
		} else if data[a] == 10 {
			if negatif == true {
				num = num * (-1)
			}
			table = append(table, num)
			num = 0
			negatif = false
		} else if data[a] == '-' {
			negatif = true
		} else {
			verif = false
			return
		}
	}
	return table
}

// Moyenne
func Average(table []float64) (res int) {
	var temp float64
	for a := 0; a < len(table); a++ {
		temp = temp + table[a]
	}
	res = int(math.Round(temp / float64(len(table))))
	return res
}

// Médiane
func Median(table []float64) (res int) {
	for a := 0; a < len(table); a++ {
		for b := 0; b < len(table)-1; b++ {
			if table[b] > table[b+1] {
				table[b], table[b+1] = table[b+1], table[b]
			}
		}
	}
	if len(table)%2 == 0 { // taille pair
		var temp float64
		milieu := (len(table) - 1) / 2
		temp = table[milieu]
		temp = temp + table[milieu+1]
		res = int(math.Round(temp / 2))
	} else { // taille impair
		milieu := len(table) / 2
		res = int(table[milieu])
	}
	return res
}

// Écart
func Variance(table []float64) (res int) {
	// Moyenne
	var temp float64
	for a := 0; a < len(table); a++ {
		temp = temp + table[a]
	}
	Moyenne := temp / float64(len(table))
	// Ecart
	var ecart_eleve float64
	for a := 0; a < len(table); a++ {
		ecart_eleve = ecart_eleve + ((table[a] - Moyenne) * (table[a] - Moyenne))
	}
	res = int(math.Round(ecart_eleve / float64(len(table))))
	return res
}

// Écart-type
func Standard_Deviation(table []float64) (res int) {
	// Moyenne
	var temp float64
	for a := 0; a < len(table); a++ {
		temp = temp + table[a]
	}
	Moyenne := temp / float64(len(table))
	// Ecart
	var ecart_eleve float64
	for a := 0; a < len(table); a++ {
		ecart_eleve = ecart_eleve + ((table[a] - Moyenne) * (table[a] - Moyenne))
	}
	ecart := ecart_eleve / float64(len(table))
	// Écart-type
	res = int(math.Round(math.Sqrt(float64(ecart))))
	return res
}

func print(s int) {
	var temp []string
	for s > 0 {
		temp = append(temp, string(rune(s%10)+'0'))
		s = s / 10
	}
	for a := len(temp) - 1; a >= 0; a-- {
		fmt.Printf(temp[a])
	}
	fmt.Println()
}
