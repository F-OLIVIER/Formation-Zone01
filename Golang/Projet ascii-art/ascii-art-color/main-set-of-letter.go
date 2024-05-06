package main

import (
	functions "ascii-art-color/packages"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Command : go run . --color=red t toto

func main() {

	// gestion des tailles d'entrée
	if len(os.Args) == 1 {
		fmt.Println("Vous devez indiquer au minimum la chaîne de caractère à traduire en ASCII")
		return
	}

	police := "standard"

	lastArg := os.Args[len(os.Args)-1]
	if lastArg == "standard" || lastArg == "thinkertoy" || lastArg == "shadow" {
		police = lastArg
		os.Args = os.Args[:len(os.Args)-1]
	}

	colors := [][]string{}

	for i := (len(os.Args) - 2); i > 0; i-- { // # NOTE : faire la boucle dans l'autre sens pour des raison de logique des messages d'erreur

		checkRes := functions.CheckParamAndReturnColor(os.Args[i])

		switch checkRes {
		case "invalidFormat":
			{
				if i == 0 {
					functions.FormatError()
					return
				} else {
					checkPreviousRes := functions.CheckParamAndReturnColor(os.Args[i-1])

					switch checkPreviousRes {
					case "invalidFormat":
						{
							functions.FormatError()
							return
						}
					case "invalidColor":
						{
							fmt.Println("La couleur du paramètre n°" + strconv.Itoa(i-1) + " est invalide")
							return
						}
					default:
						{

							for _, cArr := range colors {
								if cArr[0] == checkPreviousRes {
									fmt.Println("La couleur du paramètre n°" + strconv.Itoa(i+1) + " a déjà été saisie")
									return
								} else {

									if (strings.Contains(os.Args[i], cArr[1]) || strings.Contains(cArr[1], os.Args[i])) && len(os.Args[i]) != 0 && len(cArr[1]) != 0 {
										fmt.Println("Un chevauchement d'arguments a été détecté entre le paramètre n°" + strconv.Itoa(i) + " et le paramètre n°" + cArr[2])
										return
									}
								}
							}

							colors = append(colors, []string{checkPreviousRes, os.Args[i], strconv.Itoa(i)})
							i--
						}
					}
				}
			}
		case "invalidColor":
			{
				fmt.Println("La couleur indiquée à l'argument n°" + strconv.Itoa(i) + " est invalide")
				return
			}
		default:
			{
				for _, cArr := range colors {
					if cArr[0] == checkRes {
						fmt.Println("La couleur du paramètre n°" + strconv.Itoa(i+1) + " a déjà été saisie")
						return
					} else if cArr[1] == "" {
						fmt.Println("Il ne peux pas y avoir plusieurs couleurs de remplissage totale de la string (paramètre de couleur sans filtre)")
						return
					}
				}

				colors = append(colors, []string{checkRes, "", ""})
			}
		}

	}

	// création des couleurs

	// option d'Ascii art
	asciiSample, err := ioutil.ReadFile("./" + police + ".txt")

	if err != nil {
		fmt.Println("Une erreur est survenu lors de la lecture du fichier texte")
		return
	}

	// split les mot dans un tableau pour gérer les retour a la ligne
	asciiLines := strings.Split(strings.ReplaceAll(string(asciiSample), "\r", ""), "\n")
	asciiResult := [][]string{}

	// Boucle qui gére l'ajout de couleur et la validité des caractére saisie

	finishToColorWord := []string{"", ""}

	for index, char := range os.Args[len(os.Args)-1] { // contient le mot a print en Ascii-Art

		indexOfChar := int(char) - 32

		// len(os.Args[len(os.Args)-1])-1 correspond a la taille du dernier mot saisie
		// gestion du \n

		if char == '\\' && index != (len(os.Args[len(os.Args)-1])-2) && []rune(os.Args[len(os.Args)-1])[index+1] == 'n' { // pourquoi -2
			asciiResult = append(asciiResult, []string{"goToLine"})
		} else if indexOfChar < 0 || indexOfChar > 126 {
			// si caractére valise on l'affiche sinon on le saute et on le print pas
			continue
		}

		// gestion de l'absente ou non de filtre a colorer
		// strings.Contains recherche un match de ou des lettre a coloré
		if len(colors) != 0 {

			// si un filtre present
			res := functions.GetAsciiChar(asciiLines, indexOfChar)

			// range d'une array donc on parcour chaque ligne a coloré
			for _index, item := range res {
				for _, colorArr := range colors {
					switch true {
					case finishToColorWord[0] != "" && functions.First(strconv.Atoi(finishToColorWord[0])) >= index:
						{
							res[_index] = finishToColorWord[1] + item + "\033[0m"
						}
					case colorArr[1] != "" && index+len(colorArr[1]) <= len(os.Args[len(os.Args)-1]) && os.Args[len(os.Args)-1][index:index+len(colorArr[1])] == colorArr[1]:
						{
							finishToColorWord[0] = strconv.Itoa(index + (len(colorArr[1]) - 1))
							finishToColorWord[1] = colorArr[0]
							res[_index] = colorArr[0] + item + "\033[0m"
						}
					case colorArr[1] == "":
						{
							res[_index] = colorArr[0] + item + "\033[0m"
						}
					}

				}
			}

			asciiResult = append(asciiResult, res)

		} else {

			asciiResult = append(asciiResult, functions.GetAsciiChar(asciiLines, indexOfChar))
		}

	}

	startAt := 0

	var recursive func()

	canPrintDollar := false

	// Print de l'ascii Art

	recursive = func() {
		for i := 0; i <= 8; i++ {
			for _i, item := range asciiResult {
				if _i >= startAt {
					if item[0] == "goToLine" {
						if i == 8 {
							startAt = _i + 3
							fmt.Println("") // # NOTE : anciennement $\n
							canPrintDollar = false
							recursive()
						}
						break
					}

					fmt.Print(item[i])

					canPrintDollar = true
				}
			}
			if canPrintDollar && i != 0 {
				canPrintDollar = false
				fmt.Println("") // # NOTE : anciennement $\n
			}
		}
	}

	recursive()

}
