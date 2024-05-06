package push_swap

import "fmt"

var listeOperation []string

// fonction Chichi juste pour pas faire un gros pavet :D
func RetourDePileAetPileB(pileA, pileB []int) {
	triInferieurACent(pileA, pileB)
}

// fonction lorsque la len de la pileA est inferieur a 100 (go run ./cmd/. "2 1 3 6 5 8")
func triInferieurACent(pileA, pileB []int) {
	// première boucle for pour mettre tout les chiffres ou nombre dans la pileB et dans l'ordre décroissant
	for {
		Num := MinInt(pileA)
		// fmt.Println(" minimum :", Num)
		if len(pileA) != 3 {
			if Num != pileA[0] {
				// fmt.Println("entrer 1", pileA, pileB)
				pileA, listeOperation = OperationOnA(pileA, listeOperation)
				// fmt.Println("1", pileA, pileB)
				// fmt.Println("ok", Num)
			} else if Num == pileA[0] {
				// fmt.Println("entrer 2", pileA, pileB)
				pileA, pileB = PB(pileA, pileB)
				listeOperation = append(listeOperation, "PB")
				// fmt.Println("2", pileA, pileB)
				// fmt.Println("ok1", Num)
			} else {
				// fmt.Println("entrer dans break 1")
				// fmt.Println("ok2", Num)
				break
			}
		} else if len(pileA) == 3 && Num != pileA[0] {
			pileA, listeOperation = RepartitionPileA(pileA, listeOperation)
			// fmt.Println("4", pileA, pileB)
			// fmt.Println("ok3", Num)
		} else if len(pileA) == 3 && Num == pileA[0] {
			// fmt.Println("entrer dans break 2")
			// fmt.Println("ok4", Num)
			break
		}
	}
	// deuxièmre boucle for pour mettre tout les chiffres ou nombre contenu en pileB dans la pileA
	for {
		if len(pileB) != 0 {
			pileA, pileB = PA(pileA, pileB)
			listeOperation = append(listeOperation, "PA")
		} else {
			// fmt.Println("entrer dans break 3")
			break
		}
	}
	// impression des differentes  opération et des piles a la fin du tri
	// fmt.Printf("\nPile A : %v\t Pile B : %v\n", pileA, pileB)
	// fmt.Println("\nListe des opérations déjà effectué : ", listeOperation)
	// fmt.Println("Nombre d'opération : ", len(listeOperation))

	for _, val := range listeOperation {
		fmt.Println(val)
	}
	fmt.Println("Nombre d'opération :", len(listeOperation))

}

func OperationOnAandB(pileA, pileB []int, listeOperation []string) ([]int, []int, []string) {
	// --------------- Traitement Pile A + pile B ---------------
	// fmt.Println("-------- HERE ----------")
	if pileA[0] != MinInt(pileA) {
		if pileA[len(pileA)-1] > pileA[0] { // RRA + RRB : Ax > A1 et Bx > B1
			pileA, pileB = RRR(pileA, pileB)
			listeOperation = append(listeOperation, "RRR")
			return pileA, pileB, listeOperation
		} else if pileA[1] < pileA[0] { // SA + SB : A2 > A1 et B2 < B1
			pileA, pileB = SS(pileA, pileB)
			listeOperation = append(listeOperation, "SS")
			return pileA, pileB, listeOperation
		} else if pileA[0] > pileA[len(pileA)-1] && pileB[0] < pileB[len(pileB)-1] { // RA + RB : A1 > Ax et B1 < Bx
			pileA, pileB = RR(pileA, pileB)
			listeOperation = append(listeOperation, "RR")
			return pileA, pileB, listeOperation
		}
	}
	// fmt.Println("-------- HERE ----------")
	return pileA, pileB, listeOperation
}

func OperationOnA(pileA []int, listeOperation []string) ([]int, []string) {
	// fmt.Println("entrer OperationOnA")
	Num := MinInt(pileA)
	if len(pileA) > 1 {
		// --------------- Traitement Pile A ---------------
		if pileA[len(pileA)-1] < pileA[0] { // RRA : met Ax (dernier élément de la pile A) en haut de la pile A

			pileA = RRA(pileA)
			listeOperation = append(listeOperation, "RRA")
			return pileA, listeOperation
		} else if pileA[1] < pileA[0] { // SA : A2, A1 = A1, A2

			pileA = SA(pileA)
			listeOperation = append(listeOperation, "SA")
			return pileA, listeOperation
		} else if pileA[0] > pileA[len(pileA)-1] { // RA : met A1 en bas de la pile A

			pileA = RA(pileA)
			listeOperation = append(listeOperation, "RA")
			return pileA, listeOperation
		} else if pileA[0] > Num {
			// fmt.Println("WARNING")
			pileA = RRA(pileA)
			listeOperation = append(listeOperation, "RRA")
		}
	}
	// fmt.Println("\nListe des opérations déjà effectué : ", listeOperation)
	return pileA, listeOperation
}

func OperationOnB(pileB []int, listeOperation []string) ([]int, []string) {
	// --------------- Traitement pile B ---------------
	if len(pileB) > 1 {
		if pileB[len(pileB)-1] > pileB[0] { // RRB : met Bx (dernier élément de la pile B) en haut de la pile B1
			pileB = RRB(pileB)
			listeOperation = append(listeOperation, "RRB")
			return pileB, listeOperation
		} else if pileB[1] > pileB[0] { // SB : B2, B1 = B1, B2
			pileB = SB(pileB)
			listeOperation = append(listeOperation, "SB")
			return pileB, listeOperation
		} else if pileB[0] < pileB[len(pileB)-1] { // RB : met B1 en bas de la pile B1
			pileB = RB(pileB)
			listeOperation = append(listeOperation, "RB")
			return pileB, listeOperation
		}
	}
	// fmt.Println("\nListe des opérations déjà effectué : ", listeOperation)
	return pileB, listeOperation
}
