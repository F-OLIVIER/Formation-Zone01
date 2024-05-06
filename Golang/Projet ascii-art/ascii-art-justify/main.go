package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ASCIIChar struct {
	Ascii     string
	Caractere string
	Height    int
}

// command : go run . -align=justify 'LO th' standard

func main() {

	// récupération du mot à print

	var police string
	var arg string
	var arg1 string

	if len(os.Args) == 2 {
		// Appel le fichier police
		arg = os.Args[1]
		police = "standard"
	} else if len(os.Args) > 2 {
		// Récupére la police d'écritute (dernier arguments)
		police = os.Args[len(os.Args)-1]
		arg = os.Args[len(os.Args)-2]
		// Récupération de l'arguments de l'option de centrage
		arg1 = os.Args[1]
		// Gestion d'argument si --align existe sinon retour erreur
		flag.StringVar(&arg1, "align", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard", "A string flag")
		flag.Parse()

	} else {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}

	// Gestion de la taille du terminal
	_, width := consoleSize()
	// fmt.Println("width : ", width)

	// Variable de gestion des retours à la ligne (\n)
	var LineBreak []int

	// Tableau avec les mots à print
	arrayText := MakeArrayWithArg(arg, &LineBreak)
	count_emplacement_arrayText := 0

	// Appel le fichier police
	asciiChars := findASCIIChars(police + ".txt")

	// cherche la ligne avec le caractére
	asciiArray := findAsciiForArray(arrayText, asciiChars)
	start := 0 // setup de la variable start

	// Boucle qui met chaque mot dans une array ()
	// LineBreak correspondant au endroit ou \n est present (gestion des retours à la ligne)
	for i := 0; i < len(LineBreak)+1; i++ {
		var endIndex int
		if i == len(LineBreak) {
			endIndex = len(asciiArray)
		} else {
			endIndex = LineBreak[i]
		}

		// Génére un tableau permettant de savoir s'il y a un espace à l'emplacement d'une lettre (gestion des colonnes)
		var ArrayVerifArg []int
		// for a := 0; a < len(arg)-1; a++ {
		for _, a := range arrayText[count_emplacement_arrayText] {
			if a != ' ' {
				ArrayVerifArg = append(ArrayVerifArg, 1)
			} else {
				ArrayVerifArg = append(ArrayVerifArg, 0)
			}
		}

		// Comptage nombre d'espace dans le(s) mot(s) et récupération de l'emplacement du dernier espace
		nbespace := 0
		Emplacement_dernier_espace := 0
		temp := arrayText[count_emplacement_arrayText]
		// for i := 0; i < len(temp); i++ {
		for b, a := range temp {
			if a == ' ' {
				nbespace++
				Emplacement_dernier_espace = b
			}
		}
		// gestion des caracéres entre chaque \n
		if len(asciiArray[start:endIndex]) > 0 {
			// Fonction separateLine ==> gestion des \n
			LineToPrint := SeparateLines(strings.Join(asciiArray[start:endIndex], ""), LineBreak) //separation des lignes du mot a imprimer

			// Boucle qui print le mot
			for j := 0; j < 8; j++ {
				if arg1 == "right" {
					fmt.Println(alignRight(PrintLine(LineToPrint, j), width-len(PrintLine(LineToPrint, j)), " "))
				} else if arg1 == "left" {
					fmt.Println(PrintLine(LineToPrint, j))
				} else if arg1 == "center" {
					fmt.Println(alignCenter(width, LineToPrint, j))
				} else if arg1 == "justify" {
					if nbespace == 0 {
						// si 1 seul mot, le mettre au centre de la console
						fmt.Println(alignCenter(width, LineToPrint, j))
						continue
					}

					// taille des espaces à généré avant, entre et aprés chaque mot
					taille := (width - (len(PrintLine(LineToPrint, j)))) / nbespace
					Espace_restant := (width - (len(PrintLine(LineToPrint, j)))) % nbespace
					// fmt.Println("Espace_restant : ", Espace_restant)

					// Nombre d'espace dans l'ascii Art (espace transformé)
					// on génére une string avec le nombre d'espace souhaité si la ligne n'ai pas vide
					Str_Espace := ""
					for i := 0; i <= taille; i++ {
						Str_Espace = Str_Espace + " "
					}
					// Correction lié à la taille (1 espace + espace restant récupéré via le modulo)
					Str_Espace_restant := ""
					if Espace_restant > 0 {
						for i := 0; i < Espace_restant; i++ {
							Str_Espace_restant = Str_Espace_restant + " "
						}
					}
					Str_Espace = Str_Espace + "      "

					// remplacement des espaces par le nouveau nombre d'espace (si la colonne et un espace)
					for i := 0; i < len(LineToPrint); i++ {
						if LineToPrint[i] == "       " && ArrayVerifArg[i/8] == 0 {
							if Espace_restant > 0 && i/8 == Emplacement_dernier_espace {
								LineToPrint[i] = Str_Espace + Str_Espace_restant
							} else {
								LineToPrint[i] = Str_Espace
							}
						}
					}
					// Print de l'ascii art
					fmt.Println(PrintLine(LineToPrint, j))
				} else { // absence de --align
					fmt.Println(PrintLine(LineToPrint, j))
				}
			}
		}
		start = endIndex // nouvelle position du start pour le prochain mot
		if i != len(LineBreak)-2 && len(LineBreak) >= 2 {
			if i < len(LineBreak) && LineBreak[i-1] == LineBreak[i] {
				fmt.Println()
			}
		}
		count_emplacement_arrayText++
	}
	// Gestion du \n seul
	if len(LineBreak) != 0 && arg == "\\n" {
		fmt.Println()
	}
}

// Tableau avec les mots à print
func MakeArrayWithArg(tmpString string, LineBreak *[]int) []string {
	finalArray := make([]string, 0)
	charBuffer := ""
	dontDoNext := false
	breakLinePosstion := 0 //setup de la variable qui sert a dire la position du saut de ligne
	for i := 0; i < len(tmpString); i++ {
		if tmpString[i] == 92 && tmpString[i+1] == 110 { // Caractére \ et n
			finalArray = append(finalArray, charBuffer)
			charBuffer = ""
			dontDoNext = !dontDoNext
			*LineBreak = append(*LineBreak, breakLinePosstion) // append dans le tableau lineBreak
		} else {
			if dontDoNext {
				dontDoNext = !dontDoNext
			} else {
				breakLinePosstion++ // incrémentation quand il y a un caractère qui n'est pas un \n
				charBuffer += string(tmpString[i])
			}
		}
		if i == len(tmpString)-1 && charBuffer != "" {
			finalArray = append(finalArray, charBuffer)
		}
	}
	return finalArray
}

// Appel du fichier contenant la police
func findASCIIChars(filename string) []ASCIIChar {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}
	lines := strings.Split(string(data), "\n")
	asciiChars := make([]ASCIIChar, 0)
	charBuffer := ""
	height := 0
	foundStart := false
	asciiNumber := 32
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		if countLeadingSpaces(line) > 0 {
			if !foundStart {
				foundStart = true
			}
			charBuffer += line + "\n"
			height++
		} else if foundStart && countLeadingSpaces(charBuffer) > 0 {
			asciiChar := ASCIIChar{
				Ascii:     charBuffer,
				Caractere: string(rune(asciiNumber)),
				Height:    height,
			}
			asciiNumber++
			asciiChars = append(asciiChars, asciiChar)
			charBuffer = ""
			height = 0
			foundStart = false
		}
	}
	return asciiChars
}

// Comptage du nombre d'espace
func countLeadingSpaces(line string) int {
	return len(line) + len(strings.TrimLeft(line, " "))
}

// Cherche la ligne avec le caractére et retourne la lettre dans un tableau
func findAsciiForArray(arrayText []string, asciiChars []ASCIIChar) []string {
	var asciiArray []string
	for i := 0; i < len(arrayText); i++ {
		for k := 0; k < len(arrayText[i]); k++ {
			asciiArray = append(asciiArray, asciiChars[arrayText[i][k]-32].Ascii)
		}
	}
	return asciiArray
}

// Récupération des lettres
func SeparateLines(str string, LineBreak []int) []string {
	var sepTab []string
	tempStr := " "
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			sepTab = append(sepTab, tempStr)
			tempStr = " "
		} else {
			tempStr += string(str[i])
		}
	}
	return sepTab
}

// Fonction qui print la ligne
func PrintLine(tab []string, i int) string {
	var strBuilder strings.Builder
	for i < len(tab) {
		strBuilder.WriteString(tab[i])
		i += 8
	}
	return strBuilder.String()
}

// Récupére la taille du terminal
func consoleSize() (heigth int, width int) {
	// Création d'une commande qui demande au terminal ces dimensions
	cmd := exec.Command("stty", "size")
	// Execution de la commande dans le terminal
	cmd.Stdin = os.Stdin
	// Récupération des dimensions via out en string
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Extraction des dimensions et transformation en int
	dimension := string(out)
	espace := false
	for _, i := range dimension {
		if espace == false && i >= '0' && i <= '9' {
			heigth = heigth*10 + int(i-'0')
		} else if i == ' ' { // séparateur entre heigth et width
			espace = true
		} else if espace == true && i >= '0' && i <= '9' {
			width = width*10 + int(i-'0')
		}
	}
	return heigth, width
}

// Align a droite l'Ascii-Art
func alignRight(s string, n int, fill string) string {
	return strings.Repeat(fill, n) + s

}

// Centre l'Ascii-Art
func alignCenter(width int, LineToPrint []string, j int) (res string) {
	nb_espace_avant := (width - len(PrintLine(LineToPrint, j))) / 2
	espace_avant := ""
	for nb_espace_avant >= 0 {
		espace_avant += " "
		nb_espace_avant -= 1
	}
	res = espace_avant + PrintLine(LineToPrint, j) + espace_avant
	return res
}
