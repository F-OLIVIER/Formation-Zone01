package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Arg = []string{}

// Remplacez d'abord tous les faux sauts de ligne pour les changer en ASCII "\n",
// cela coupera ensuite l'entrée en sous-chaînes séparées par "\n"
// et renverra un tableau de chaînes contenant toutes les sous-chaînes.
// Appeler dans ./main/main.go
func AsciiArt(s string) []string {
	// remplacement des faux saut de ligne \\n par \n
	s = strings.ReplaceAll(s, "\\n", "\n")
	str := []rune(s)
	arr := []string{}
	tmp := ""

	// gestion des \n en retour a la ligne et non en impression ascii-art de "\" et "n"
	for i := 0; i < len(str); i++ {
		ru := str[i]
		switch {
		case ru == '\n' && tmp == "":
			arr = append(arr, "\n")
		case ru == '\n':
			arr = append(arr, tmp)
			tmp = ""
		default:
			tmp += string(ru)
		}
	}

	if tmp != "" {
		arr = append(arr, tmp)
	}

	return arr
}

// Imprimez chaque lettre sous forme de caractère ASCII dans la console en récupérant l'art
// dans une base de données .txt dans un dossier.
// Appeler dans ./main/main.go
func PrintWords(arr []string) { // prend en entrée l'array de string du résultats d'Ascii Art
	line := ""

	for _, word := range arr {
		for i := 0; i < 8; i++ {
			for _, ru := range word {
				if ru != '\n' {
					letter := i + (int(ru)-32)*9 + 2
					line += findLetter(letter)
				}
			}
			if word != "\n" || i == 1 {
				fmt.Println(line)
				line = ""
			}
		}
	}
}

// fonction utilisée pour récupérer chaque lettre de la base de données du fichier.
// Appeler dans TestAsciiArt et utilisé dans PrintWords (ici)
func findLetter(lineNum int) string {
	file, err := os.Open("standard.txt")
	if err != nil {
		file, _ = os.Open("main/standard.txt")
	}

	sc := bufio.NewScanner(file)
	lastLine := 0

	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			tmp := sc.Text()
			file.Close()
			return tmp
		}
	}

	return ""
}

// Cette fonction agit comme PrintWord() sauf qu'elle renvoie la valeur sous forme de chaîne occupant
// la première position d'un tableau de chaînes.
// Appeler dans ascii-art_test.go uniquement
func TestAsciiArt(arr []string) []string {
	final := []string{}
	line := ""

	for _, word := range arr {
		for i := 0; i < 8; i++ {
			for _, ru := range word {
				if ru != '\n' {
					letter := i + (int(ru)-32)*9 + 2
					line += findLetter(letter)
				}
			}

			if word != "\n" || i == 1 {
				line += "\n"
			}
		}

	}

	final = append(final, line)

	return final
}
