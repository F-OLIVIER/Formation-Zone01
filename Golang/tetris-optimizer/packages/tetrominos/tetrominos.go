package tetris

import (
	"fmt"
	"os"
)

var Listetetromino [19]([4][4]string)
var Tetrominosvalider = [19]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var ArrayFinal [][]string
var lettreAPlacer = 65
var color = [19]string{"\x1b[48;2;220;55;20m", "\x1b[48;2;128;255;0m", "\x1b[48;2;0;255;243m", "\x1b[48;2;0;50;255m", "\x1b[48;2;255;255;255m", "\x1b[48;2;228;0;255m", "\x1b[48;2;255;0;4m", "\x1b[48;2;155;0;255m", "\x1b[48;2;255;112;0m", "\x1b[48;2;255;247;0m", "\x1b[48;2;255;0;240m", "\x1b[48;2;0;209;3m", "\x1b[48;2;0;255;236m", "\x1b[48;2;203;203;203m", "\x1b[48;2;197;117;234m", "\x1b[48;2;240;79;240m", "\x1b[48;2;255;214;106m", "\x1b[48;2;15;0;255m", "\x1b[48;2;155;106;106m"}

func Listetetrominos() {
	// liste des 19 tétrominos possible (7 tetrominos dans toutes les positions possible)
	Listetetromino[0] = [4][4]string{{" ", "#", " ", " "}, {" ", "#", " ", " "}, {"#", "#", " ", " "}, {" ", " ", " ", " "}}  // tetromino 1
	Listetetromino[1] = [4][4]string{{"#", "#", " ", " "}, {" ", "#", " ", " "}, {" ", "#", " ", " "}, {" ", " ", " ", " "}}  // tetromino 2
	Listetetromino[2] = [4][4]string{{"#", "#", "#", " "}, {" ", " ", "#", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 3
	Listetetromino[3] = [4][4]string{{" ", " ", " ", "#"}, {"#", "#", "#", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 4
	Listetetromino[4] = [4][4]string{{"#", "#", " ", " "}, {"#", " ", " ", " "}, {"#", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 5
	Listetetromino[5] = [4][4]string{{"#", " ", " ", " "}, {"#", " ", " ", " "}, {"#", "#", " ", " "}, {" ", " ", " ", " "}}  // tetromino 6
	Listetetromino[6] = [4][4]string{{"#", " ", " ", " "}, {"#", "#", "#", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 7
	Listetetromino[7] = [4][4]string{{"#", "#", "#", " "}, {"#", " ", " ", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 8
	Listetetromino[8] = [4][4]string{{" ", "#", "#", " "}, {"#", "#", " ", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 9
	Listetetromino[9] = [4][4]string{{"#", "#", " ", " "}, {" ", "#", "#", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}}  // tetromino 10
	Listetetromino[10] = [4][4]string{{"#", " ", " ", " "}, {"#", "#", " ", " "}, {" ", "#", " ", " "}, {" ", " ", " ", " "}} // tetromino 11
	Listetetromino[11] = [4][4]string{{" ", "#", " ", " "}, {"#", "#", " ", " "}, {"#", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 12
	Listetetromino[12] = [4][4]string{{"#", "#", "#", " "}, {" ", "#", " ", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 13
	Listetetromino[13] = [4][4]string{{" ", "#", " ", " "}, {"#", "#", "#", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 14
	Listetetromino[14] = [4][4]string{{"#", " ", " ", " "}, {"#", "#", " ", " "}, {"#", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 15
	Listetetromino[15] = [4][4]string{{" ", "#", " ", " "}, {"#", "#", " ", " "}, {" ", "#", " ", " "}, {" ", " ", " ", " "}} // tetromino 16
	Listetetromino[16] = [4][4]string{{"#", "#", "#", "#"}, {" ", " ", " ", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 17
	Listetetromino[17] = [4][4]string{{"#", " ", " ", " "}, {"#", " ", " ", " "}, {"#", " ", " ", " "}, {"#", " ", " ", " "}} // tetromino 18
	Listetetromino[18] = [4][4]string{{"#", "#", " ", " "}, {"#", "#", " ", " "}, {" ", " ", " ", " "}, {" ", " ", " ", " "}} // tetromino 19
}

// fonction pour assemblé les piéces du Tetris
func Writetetris() {
	// fmt.Println("Tetrominosvalider : ", Tetrominosvalider)
	// Definition de la taille de l'array final et création de l'array final
	arrayfinal()
	// Application de la résolution du tetris
	if Resolution() == true { // si valide impression du Tetris
		printTetris()
	} else { // si non valide, retour 'Error'
		fmt.Println("Error")
		os.Exit(1)
	}

}

// Fonction de résolution : basé sur l'algorihtmme de Backtracking
// explication compléte de l'algorihtmme pour le Sudoku : http://igm.univ-mlv.fr/~dr/XPOSE2013/sudoku/backtracking.html#:~:text=Résolution%20du%20Sudoku&text=On%20part%20de%20la%20cellule,on%20place%20la%20première%20possible.
func Resolution() bool {
	var ligne, colonne int
	// Fonction qui permet de sortir de la récursive : vérifie si toutes les valeurs sont saisie dans le Sudoku, c'est à dire plus aucune case vide
	terminer, tetroAPlacer := tetrominoNonAssignée()
	if terminer == true {
		return true
	}
	tetrominosValideraPlacer := Listetetromino[tetroAPlacer]

	// Récupération des cases non assigné
	coordonnees := caseNonAssignée(&ligne, &colonne)

	// Recherche d'une case vide
	for _, a := range coordonnees {
		ligne = a[0]
		colonne = a[1]
		if emplacementTetroValide(&ligne, &colonne, &tetrominosValideraPlacer) == true {
			// fmt.Println("emplacementTetroValide : True")
			// Placement du tetromino
			placementTetro(&ligne, &colonne, &tetrominosValideraPlacer, &tetroAPlacer)
			// fmt.Println("Tetris en cour : ")
			// printTetris()
			// si la résolution est terminé on renvoie "true pour cloturer la récursivité
			if Resolution() == true {
				return true
			}
			// reinitiaisation les emplacements du tetrominos avant le changement des coordonnées x et y
			retraitTetro(&ligne, &colonne, &tetrominosValideraPlacer, &tetroAPlacer)
		}
	}
	// }
	return false
}

// fonction qui renvoie les corordonées x et y de la prochaine case vide
func tetrominoNonAssignée() (bool, int) {
	tetroAPlacer := 0
	for i := 0; i < 19; i++ {
		if Tetrominosvalider[i] != 0 {
			return false, i
		}
	}
	return true, tetroAPlacer
}

// fonction qui renvoie les corordonées x et y de la prochaine case vide
func caseNonAssignée(ligne *int, colonne *int) []([]int) {
	var coordonnees []([]int)
	for i := 0; i < len(ArrayFinal); i++ {
		for j := 0; j < len(ArrayFinal); j++ {
			if ArrayFinal[i][j] == "." {
				*ligne = i
				*colonne = j
				// fmt.Println("caseNonAssignée : ", *ligne, *colonne)
				tmp := []int{i, j}
				coordonnees = append(coordonnees, tmp)
			}
		}
	}
	return coordonnees
}

// test si les 3 fonctions de validation (ligne + colonne + case) sont ok
func emplacementTetroValide(ligne *int, colonne *int, tetrominosValideraPlacer *[4][4]string) bool {
	// fmt.Println("tetrominosValideraPlacer : ", *tetrominosValideraPlacer)
	var coordonneeTetro []([]int)
	// boucle qui récupére les coordonnées des emplacements du tretrominos à placer
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetrominosValideraPlacer[i][j] == "#" {
				tmp := []int{i, j}
				coordonneeTetro = append(coordonneeTetro, tmp)
			}
		}
	}

	// fmt.Println("coordonneeTetro : ", coordonneeTetro)
	// Check des coordonnées récupérer (si "." partout, le tetromino peux etre mis)
	for _, i := range coordonneeTetro {
		// Si l'emplacement sort du carré
		if ((i[0] + *ligne) > len(ArrayFinal)-1) || ((i[1] + *colonne) > len(ArrayFinal)-1) {
			return false
		}
		// Si l'emplacement a déja un tetromino placé
		if ArrayFinal[i[0]+*ligne][i[1]+*colonne] != "." {
			return false
		}
	}
	return true
}

// fonction qui place un tetrominos au coordonées x et y données
func placementTetro(ligne *int, colonne *int, tetrominosValideraPlacer *[4][4]string, tetroAPlacer *int) {
	var coordonneeTetro []([]int)
	// boucle qui récupére les coordonnées des emplacements du tretrominos à placer
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetrominosValideraPlacer[i][j] == "#" {
				tmp := []int{i, j}
				coordonneeTetro = append(coordonneeTetro, tmp)
			}
		}
	}

	// Retrait du tetrominos placé
	Tetrominosvalider[*tetroAPlacer] = Tetrominosvalider[*tetroAPlacer] - 1
	// Placement du tetrominos
	for _, i := range coordonneeTetro {
		ArrayFinal[i[0]+*ligne][i[1]+*colonne] = color[*tetroAPlacer] + string(rune(lettreAPlacer)) + "\x1b[0m"
	}
	// fmt.Println("tetromino placé")
	lettreAPlacer++
}

// fonction qui retire un tetrominos au coordonées x et y données
func retraitTetro(ligne *int, colonne *int, tetrominosValideraPlacer *[4][4]string, tetroAPlacer *int) {
	var coordonneeTetro []([]int)
	// boucle qui récupére les coordonnées des emplacements du tretrominos à placer
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetrominosValideraPlacer[i][j] == "#" {
				tmp := []int{i, j}
				coordonneeTetro = append(coordonneeTetro, tmp)
			}
		}
	}
	// Remise du tetrominos placé
	Tetrominosvalider[*tetroAPlacer] = Tetrominosvalider[*tetroAPlacer] + 1
	// Retrait du tetrominos
	for _, i := range coordonneeTetro {
		ArrayFinal[i[0]+*ligne][i[1]+*colonne] = "."
	}

	lettreAPlacer--
}

// Création du tableau d'affichage final
func arrayfinal() {
	nbtetromino := 0
	// Calcul du nombre de case necessaire
	for _, i := range Tetrominosvalider {
		nbtetromino = nbtetromino + i
	}
	taillearrayfinal := 4 * nbtetromino
	// Gestion d'erreur si grille de base contien les 2 barres qui necessite 5 cases au lieu de 4
	if taillearrayfinal <= 16 {
		if Tetrominosvalider[16] == 1 && Tetrominosvalider[17] == 1 { // Si tetromino de taille 4 hauteur et 4 largeur
			taillearrayfinal = 4 * 5
		} else if Tetrominosvalider[16] == 1 || Tetrominosvalider[17] == 1 { // Si tetromino de taille 4 (hauteur ou largeur)
			taillearrayfinal = 4 * 4
		} else if Tetrominosvalider[1] == 1 || // Si tetromino de taille 3 (hauteur ou largeur)
			Tetrominosvalider[2] == 1 ||
			Tetrominosvalider[3] == 1 ||
			Tetrominosvalider[4] == 1 ||
			Tetrominosvalider[5] == 1 ||
			Tetrominosvalider[6] == 1 ||
			Tetrominosvalider[7] == 1 ||
			Tetrominosvalider[8] == 1 ||
			Tetrominosvalider[9] == 1 ||
			Tetrominosvalider[10] == 1 ||
			Tetrominosvalider[11] == 1 ||
			Tetrominosvalider[12] == 1 ||
			Tetrominosvalider[13] == 1 ||
			Tetrominosvalider[14] == 1 ||
			Tetrominosvalider[15] == 1 ||
			Tetrominosvalider[16] == 1 {
			taillearrayfinal = 4 * 2
		}
	}
	// Boucle pour définir le nombre de case minimum a avoir et création de l'array
	for i := 2; i <= taillearrayfinal; i++ {
		if (i * i) >= taillearrayfinal {
			ArrayFinal = make([][]string, i)
			for j := 0; j < i; j++ {
				for k := 0; k < i; k++ {
					ArrayFinal[j] = append(ArrayFinal[j], ".")
				}
			}
			return
		}
	}
}

func printTetris() {
	vide := 0
	result := ""
	for _, i := range ArrayFinal {
		for _, j := range i {
			result = result + j
			if j == "." {
				vide++
			}
		}
		result = result + "\n"
	}
	fmt.Println()
	fmt.Println("Tetris terminé avec", vide, "emplacement(s) vide")
	fmt.Println(result)
}
