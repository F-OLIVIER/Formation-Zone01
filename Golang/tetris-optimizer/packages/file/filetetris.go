package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	TETRIS "tetris-optimizer/packages/tetrominos"
)

// fonction pour lire une ligne d'un fichier
func ReadLineFile(namefile string) {
	// Ouverture du fichier
	file, err := os.Open("./Fichier_Tetrominos/" + namefile)
	if err != nil {
		fmt.Printf("ERROR, probléme avec le nom du fichier : ")
		fmt.Println(err)
		os.Exit(1)
	}

	// Création d'un scanner
	sc := bufio.NewScanner(file)

	// Scan du fichier
	interline := 1 // ligne intermédiaire
	tetrominos := [4][4]string{}
	// Définition de la liste des tetrominos
	TETRIS.Listetetrominos()
	var errtetris error
	for sc.Scan() { // ligne courante
		// fmt.Println(sc.Text())
		if interline%5 == 0 {
			if len(sc.Text()) != 0 {
				fmt.Println("ERROR")
				os.Exit(1)
			}
		}
		if interline%5 != 0 { // toutes les 5 lignes
			for i, j := range sc.Text() {
				// Si taille différente de 4, trop de caractére dans la ligne du fichier
				if len(sc.Text()) != 4 {
					fmt.Println("Un caractéres contenu dans le fichier ne correspondant pas à un caractére Tetris")
					os.Exit(1)
				}
				// parcourt l'arguments d'entrée qui est en ligne pour compléter un tableau 9x9
				tetrominos[interline-1][i] = string(j)
			}
			interline++
		} else {
			// fmt.Println("tetrominos : ", tetrominos)
			errtetris = checktetrominos(tetrominos)
			if errtetris != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			interline = 1
			tetrominos = [4][4]string{}
		}
	}
	// fmt.Println("tetrominos : ", tetrominos)
	errtetris = checktetrominos(tetrominos)
	if errtetris != nil {
		fmt.Println(errtetris)
		os.Exit(1)
	}
}

// fonction pour assemblé les piéces du Tetris
func checktetrominos(table [4][4]string) error {
	nbcaract := 0

	linebase := 3
	colbase := 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] == "#" {
				// pour gestion d'erreur
				nbcaract++
				// Récupération des coordonnées de base (debut caractéres en haut a gauche)
				if i < linebase {
					linebase = i
				}
				if j < colbase {
					colbase = j
				}
			}
		}
	}

	// Suppression des lignes et colonnes vide
	basetetrominos := [4][4]string{}
	colcourante := colbase
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if table[linebase][colcourante] == "." {
				basetetrominos[i][j] = " "
			} else {
				basetetrominos[i][j] = table[linebase][colcourante]
			}
			colcourante++
			if colcourante == 4 {
				break
			}
		}
		linebase++
		colcourante = colbase
		if linebase == 4 {
			break
		}
	}
	// fmt.Println("basetetrominos : ", basetetrominos)

	// Remplacement des vides par des espaces
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if basetetrominos[i][j] == "" {
				basetetrominos[i][j] = " "
			}
		}
	}

	errcompare := CompareTetrominos(basetetrominos)
	if errcompare != nil {
		return errcompare
	}

	// gestion d'erreur si trop ou pas assez de caractére dans le tetrominos
	if nbcaract != 4 {
		return errors.New("ERROR : Un caractéres contenu dans le fichier ne correspondant pas à un caractére Tetris")
	}

	return nil
}

func CompareTetrominos(basetetrominos [4][4]string) error {
	// Boucle de comparaison
	for i, j := range TETRIS.Listetetromino {
		if basetetrominos == j {
			TETRIS.Tetrominosvalider[i] = TETRIS.Tetrominosvalider[i] + 1
			return nil
		}
	}
	return errors.New("ERROR : Un caractéres contenu dans le fichier ne correspondant pas à un caractére Tetris")
}
