package main

import (
	internal "push-swap/checkers/internal"
)

func main() {
	// gestion du contenu int
	pileA, pileB := internal.VerifyAndConvert()
	// fmt.Println(pileA, pileB)

	// gestion du contenu avant le pipe
	listOperation := internal.Pipe()
	// fmt.Println("enter : ", listOperation)

	// check des opérations
	internal.Traitement(listOperation, pileA, pileB)
}
