package ascii

import (
	"log"
	"os"
	"strings"
)

// Transforme la bibliotheque standard
func Biblio(input string) []string {
	police := "template/standard.txt"
	if input != "standard" {
		switch input {
		case "shadow":
			police = "template/shadow.txt"
		case "thinkertoy":
			police = "template/thinkertoy.txt"
		case "varsity":
			police = "template/varsity.txt"
		default:
			log.Fatalln("Police", input, "Introuvable")
		}
	}

	in, _ := os.ReadFile(police)

	var str string
	for _, k := range in {
		if k != 13 {
			str += string(k)
		}
	}
	bib := strings.Split(str, "\n")
	return bib
}

// Fonction qui transforme l'argument en ascii art
func Ascii_Return(input string, arg string) string {
	// Créer un arr pour split aux \n
	arr := strings.Split(input, "\\n")

	bib := Biblio(arg)
	var ascii string
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			ascii += Ascii_Write(arr[i], bib)
		}
		if arr[i] == "" {
			ascii += "\n"
		}
	}
	return ascii
}

// Fonction qui écrit une ligne de caractère en ascii art en une string
func Ascii_Write(str string, bib []string) string {
	var ascii string
	for i := 0; i <= 7; i++ {
		for _, k := range str {
			ascii += bib[(k-32)*9+1+rune(i)]
		}
		ascii += "\n"
	}
	return ascii
}
