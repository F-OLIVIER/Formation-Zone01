package push_swap

import (
	"fmt"
	"os"
)

func Traitement(listOperation []string, pileA, pileB []int) {
	// Traitement des opérations de listOpération
	for _, operation := range listOperation {
		if operation == "PA" {
			if len(pileB) > 0 {
				pileA, pileB = PA(pileA, pileB)
			} else {
				fmt.Println("KO PA")
				os.Exit(0)
			}
		} else if operation == "PB" {
			if len(pileA) > 0 {
				pileA, pileB = PB(pileA, pileB)
			} else {
				fmt.Println("KO PB")
				os.Exit(0)
			}
		} else if operation == "SA" {
			if len(pileA) > 2 {
				pileA = SA(pileA)
			} else {
				fmt.Println("KO SA")
				os.Exit(0)
			}
		} else if operation == "SB" {
			if len(pileB) > 2 {
				pileB = SB(pileB)
			} else {
				fmt.Println("KO SB")
				os.Exit(0)
			}
		} else if operation == "SS" {
			if len(pileA) > 2 && len(pileB) > 2 {
				pileA, pileB = SS(pileA, pileB)
			} else {
				fmt.Println("KO SS")
				os.Exit(0)
			}
		} else if operation == "RA" {
			if len(pileA) > 2 {
				pileA = RA(pileA)
			} else {
				fmt.Println("KO RA")
				os.Exit(0)
			}
		} else if operation == "RB" {
			if len(pileB) > 2 {
				pileA = RB(pileB)
			} else {
				fmt.Println("KO RB")
				os.Exit(0)
			}
		} else if operation == "RR" {
			if len(pileA) > 2 && len(pileB) > 2 {
				pileA, pileB = RR(pileA, pileB)
			} else {
				fmt.Println("KO RR")
				os.Exit(0)
			}
		} else if operation == "RRA" {
			if len(pileA) > 2 {
				pileA = RRA(pileA)
			} else {
				fmt.Println("KO RRA")
				os.Exit(0)
			}
		} else if operation == "RRB" {
			if len(pileB) > 2 {
				pileB = RRB(pileB)
			} else {
				fmt.Println("KO RRB")
				os.Exit(0)
			}
		} else if operation == "RRR" {
			if len(pileA) > 2 && len(pileB) > 2 {
				pileA, pileB = RRR(pileA, pileB)
			} else {
				fmt.Println("KO RRR")
				os.Exit(0)
			}
		}
	}

	// vérification du trie de la pileA
	for i := 0; i < len(pileA)-1; i++ {
		if pileA[i] > pileA[i+1] {
			fmt.Println("KO")
			os.Exit(0)
		}
	}
	fmt.Println("OK")
}
