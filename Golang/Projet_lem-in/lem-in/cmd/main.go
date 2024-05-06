package main

import (
	"fmt"
	"os"
	"time"

	internal "lem-in/internal"
)

// command : go run . example/example00.txt

func main() {
	duree_execution_debut := time.Now()
	internal.Travail(os.Args[1:])
	// Print de la durée d'execution du programme
	duree_execution_fin := time.Now()
	internal.PrintFourmis()
	fmt.Printf("Toutes les fourmis ont traversé la fourmiliéres en %v\n", duree_execution_fin.Sub(duree_execution_debut))
}
