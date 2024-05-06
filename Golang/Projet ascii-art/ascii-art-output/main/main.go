package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"asciiartoutput"
)

// command : go run . --output=sample.txt --color=red t toto standard

func main() {

	NameOutputFile := flag.String("output", "invalid", "Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --output=sample.txt something shadow")
	color := flag.String("color", "invalid", "Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --color=red something")
	flag.Parse()

	// Récupération des arguments aprés la banner
	args := flag.Args()
	argsTesting := os.Args[1:]

	// Usage des Prefix
	var outputafter []string
	var newArgs []string
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--output=") {
			outputafter = append(outputafter, strings.TrimPrefix(args[i], "--output="))
		} else {
			newArgs = append(newArgs, args[i])
		}
	}
	args = newArgs

	// Si output non mis en 1er option d'arguments
	if len(outputafter) == 1 {
		*NameOutputFile = outputafter[0]
	}

	// gestion du nombre d'argument hors flag
	if len(args) > 3 {
		fmt.Println("Number of arguments invalid\nUsage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --output=sample.txt something")
		os.Exit(0)
	}

	var colorOutput string
	var banner string
	BannerPresente := false
	var input string
	Colour := true
	Letter := false
	Output := false
	var Colorletter string
	if color != nil && *color != "invalid" {
		colorOutput = asciiartoutput.ColorOutput(*color)
	} else {
		Colour = false
	}

	if NameOutputFile != nil && *NameOutputFile != "invalid" {
		Output = true
		// Création du fichier de sortie
		data, err := os.Create(*NameOutputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
	}

	if args[len(args)-1] == "standard" || args[len(args)-1] == "shadow" || args[len(args)-1] == "thinkertoy" { // on regarde si il y a une bannière
		banner = args[len(args)-1] + ".txt"
		input = args[len(args)-2]
		BannerPresente = true
	} else {
		banner = "standard.txt"
		input = args[len(args)-1]
	}

	// gestion du nombre d'arguments SANS l'option couleur
	if !Colour && len(args) == 1 {
		// laisser passer
	} else if !Colour && len(args) > 2 {
		if BannerPresente == false {
			fmt.Println("Too many arguments")
			os.Exit(0)
		}
	}

	// gestion du nombre d'arguments AVEC l'option couleur
	if Colour && len(args) == 1 {
		// laisser passer
	} else if Colour && len(args) > 3 {
		if BannerPresente == false {
			fmt.Println("Too many arguments")
			os.Exit(0)
		}
		Letter = true
	} else {
		Letter = true
	}

	// Gestion de la lettre a coloré (si existante)
	if (len(args) == 2 && BannerPresente == false) || (len(args) == 3 && BannerPresente == true) {
		Colorletter = args[0]
	}

	// Ouverture du fichier ascii Art
	_, err := os.Open(banner)
	if err != nil {
		fmt.Println("Can't open art file")
		fmt.Println(err)
		os.Exit(0)
	}

	// Retour erreur si pas de '=' ou '--' dans l'option --color
	if (strings.Contains(argsTesting[0], "--color") && !strings.Contains(argsTesting[0], "=")) ||
		(len(argsTesting) > 1 && strings.Contains(argsTesting[1], "--color") && !strings.Contains(argsTesting[1], "=")) {
		fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . something standard")
		os.Exit(0)
	}
	// Retour erreur si pas de '=' ou '--' dans l'option --output
	if (strings.Contains(argsTesting[0], "--output") && !strings.Contains(argsTesting[0], "=")) ||
		(len(argsTesting) > 1 && strings.Contains(argsTesting[1], "--output") && !strings.Contains(argsTesting[1], "=")) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	splittedLines := asciiartoutput.SplitLines(input)

	// si il n'y a pas de couleur , on gère l'input et la banner
	if Colour == false {
		for _, text := range splittedLines {
			tableau := asciiartoutput.Asciiart(text, banner)
			for i := 0; i < 8; i++ {
				if Output == true {
					// Ouverture du fichier
					file, _ := os.OpenFile(*NameOutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
					defer file.Close()                  // on ferme automatiquement à la fin de notre programme
					file.WriteString(tableau[i] + "\n") // écrire dans le fichier
				} else {
					fmt.Println(tableau[i])
				}
			}
		}
		if Output == true {
			fmt.Println("Contenu écrit dans le fichier " + *NameOutputFile)
		}
		return
	}

	// Presence d'une couleur et absence d'une lettre spécifique
	if Colour && !Letter {
		for _, text := range splittedLines {
			tableau := asciiartoutput.Asciiart(text, banner)
			for i := 0; i < 8; i++ {
				if Output == true {
					// Ouverture du fichier
					file, _ := os.OpenFile(*NameOutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
					defer file.Close() // on ferme automatiquement à la fin de notre programme
					file.WriteString(colorOutput + tableau[i] + asciiartoutput.Reset + "\n")
				} else {
					fmt.Println(colorOutput + tableau[i] + asciiartoutput.Reset)
				}
			}
		}
		if Output == true {
			fmt.Println("Contenu écrit dans le fichier " + *NameOutputFile)
		}
		return
	}

	forFileWritre := ""
	// Presence d'une couleur et d'une lettre spécifique
	if Colour && Letter {
		for _, str := range splittedLines {
			tabColor := asciiartoutput.SplitColor(str, Colorletter) // on sépare les lettres à colorer
			lignes := [8]string{}
			for _, s := range tabColor {
				tableau := asciiartoutput.Asciiart(s, banner)
				if s == Colorletter {
					for i := 0; i < 8; i++ {
						lignes[i] += colorOutput + tableau[i] + asciiartoutput.Reset
					}

				} else {
					for i := 0; i < 8; i++ {
						lignes[i] += tableau[i]
					}
				}

			}
			for i := 0; i < 8; i++ {
				forFileWritre = forFileWritre + lignes[i] + "\n"
			}
			if len(splittedLines) > 1 {
				forFileWritre = forFileWritre + "\n"
			}

		}
		if Output == true {
			// Ouverture du fichier
			file, _ := os.OpenFile(*NameOutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
			defer file.Close() // on ferme automatiquement à la fin de notre programme
			file.WriteString(forFileWritre)
			fmt.Println("Contenu écrit dans le fichier " + *NameOutputFile)
		} else {
			fmt.Println(forFileWritre)
		}
		return
	}

}
