package main

import (
	"fmt"
	"os"
	FILE "tetris-optimizer/packages/file"
	TETRIS "tetris-optimizer/packages/tetrominos"
	"time"
)

// Command : go run . sample.txt

func main() {
	debut := time.Now()
	// gestion d'erreurs d'entrée
	var args []string
	if len(os.Args[1:]) == 1 {
		args = os.Args[1:]
	} else {
		fmt.Println("Probleme dans le nombre d'argument saisie, la commande necessite le nom du fichier")
		fmt.Println("exemple d'utilisation : \"go run . sample.txt\"")
		os.Exit(0)
	}

	// Fonction pour lire une ligne d'un fichier
	FILE.ReadLineFile(args[0])

	// Fonction pour assemblé les piéces du Tetris
	TETRIS.Writetetris()

	fin := time.Now()
	fmt.Println("Temps d'éxécutiion : ", fin.Sub(debut))
}
