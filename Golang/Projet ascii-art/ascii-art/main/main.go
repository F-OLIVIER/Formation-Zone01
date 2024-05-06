package main

import (
	"ascii"
	"fmt"
	"os"
)

func main() {
	// Gestion de l'arguments de choix de type d'Ascii Art (standard ou thinkertoy ou shadow)
	arg := findArgs()
	// Génére l'Ascii Art
	art := ascii.AsciiArt(arg[0])
	// Imprime l'Ascii Art
	ascii.PrintWords(art)
}

// Gestion de l'arguments de choix de type d'Ascii Art (standard ou thinkertoy ou shadow)
// doublon dans ascii-art.go
func findArgs() []string {
	args := os.Args[1:]
	l := len(args)
	// Gestion d'erreur
	if l < 1 || l > 2 {
		fmt.Println("One(1) or Two(2) arguments only, please")
		os.Exit(1)
	}

	// Récupération du type d'Ascii Art
	if len(args) == 1 {
		args = append(args, "standard.txt")
	} else {
		switch args[1] {
		case "-t":
			args[1] = "thinkertoy.txt"
		case "-s":
			args[1] = "shadow.txt"
		default:
			os.Exit(2)
		}
	}

	return args
}
