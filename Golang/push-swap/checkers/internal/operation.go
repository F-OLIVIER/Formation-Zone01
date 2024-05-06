package push_swap

// function pour pousser B1 en haut de la pile A
func PA(pileA, pileB []int) ([]int, []int) {
	// on récupère le premier argument de la pile B
	pos1 := pileB[0]
	// on crée une slice de pos1
	npos1 := []int{pos1}
	// on ajoute npos1 a la pile a en première position
	pileA = append(npos1, pileA...)
	pileB = pileB[1:]
	return pileA, pileB
}

//----------------------------------------------------------------

// function pour pousser A1 en haut de la pile B
func PB(pileA, pileB []int) ([]int, []int) {
	// on récupère le premier argument de la pile A
	pos1 := pileA[0]
	// on crée une slice de pos1
	npos1 := []int{pos1}
	// on ajoute npos1 a la pile b en première position
	pileB = append(npos1, pileB...)
	pileA = pileA[1:]
	return pileA, pileB
}

//----------------------------------------------------------------

// function pour inverser A1 et A2
func SA(pileA []int) []int {
	// A2, A1 = A1, A2
	A1 := pileA[0]
	A2 := pileA[1]
	// inversion de A1 et A2
	A1, A2 = A2, A1
	// création d'une slice de A1 et A2
	npileA := []int{A1, A2}
	// integration de pilA à npileA pour mettre A1 et A2 dans la même pile et en première position
	pileA = append(npileA, pileA[2:]...)
	return pileA
}

//----------------------------------------------------------------

// function pour inverser B1 et B2
func SB(pileB []int) []int {
	// B2, B1 = B1, B2
	B1 := pileB[0]
	B2 := pileB[1]
	// inversion de B1 et B2
	B1, B2 = B2, B1
	// création d'une slice de B1 et B2
	npileB := []int{B1, B2}
	// integration de pilB à npileB pour mettre B1 et B2 dans la même pile et en première position
	pileB = append(npileB, pileB[2:]...)
	return pileB
}

//----------------------------------------------------------------

// fonction qui fait légèrement nazi :D (la blagounette du mois)et qui rassemble les functions SA et SB
func SS(pileA, pileB []int) ([]int, []int) {
	return SA(pileA), SB(pileB)
}

//----------------------------------------------------------------

// function pour mettre A1 en bas de pile A
func RA(pileA []int) []int {
	// prend le 1er argument de la pile A
	A1 := pileA[0]
	// ajoute A1 à la pile A
	pileA = append(pileA[1:], A1)
	return pileA
}

//----------------------------------------------------------------

// function pour mettre B1 en bas de pile B
func RB(pileB []int) []int {
	// prend le 1er argument de la pile B
	B1 := pileB[0]
	// ajoute B1 à la pile B
	pileB = append(pileB[1:], B1)
	return pileB
}

//----------------------------------------------------------------

// function qui regroupe les fonctions RA et RB
func RR(pileA, pileB []int) ([]int, []int) {
	return RA(pileA), RB(pileB)
}

//----------------------------------------------------------------

// function pour mettre Ax (dernier argument de la pile A) en haut de pile A
func RRA(pileA []int) []int {
	// on récupère le dernière argument de la pile A dans la variable Ax
	Ax := pileA[len(pileA)-1]
	// on met Ax dans la variable Sax qui est une slice de Ax
	Sax := []int{Ax}
	// on rajoute Sax a la pile A en supprimant le dernier élément de celle-ci
	pileA = append(Sax, pileA[:len(pileA)-1]...)
	return pileA
}

//----------------------------------------------------------------

// function pour mettre Bx (dernier argument de la pile B) en haut de pile B
func RRB(pileB []int) []int {
	// on récupère le dernière argument de la pile B dans la variable Bx
	Bx := pileB[len(pileB)-1]
	// on met Bx dans la variable Sax qui est une slice de Bx
	Sax := []int{Bx}
	// on rajoute Sax a la pile A en supprimant le dernier élément de celle-ci
	pileB = append(Sax, pileB[:len(pileB)-1]...)
	return pileB
}

//----------------------------------------------------------------

// function qui regroupe les fonctions RRA et RRB
func RRR(pileA, pileB []int) ([]int, []int) {
	return RRA(pileA), RRB(pileB)
}

//----------------------------------------------------------------
