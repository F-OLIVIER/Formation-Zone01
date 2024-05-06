package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Récupération des tailles arguments
	args := os.Args[1:]
	if len(args) > 2 || len(args) < 1 {
		printUsageAndExit()
	}

	// Récupération des arguments mot et banner avec gestion de la presence de la banner ou non
	str := args[0]
	var banner string
	if len(args) == 2 {
		if args[1] != "standard" && args[1] != "shadow" && args[1] != "thinkertoy" {
			printUsageAndExit()
		}
		banner = args[1]
	} else {
		banner = "standard"
	}
	// génére l'Ascii
	verificationRetourLigne(str, banner)
}

func printUsageAndExit() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . something standard")
	os.Exit(0)
}

// fonction pour lire une ligne d'un fichier
func readLine(lineNum int, banner string) (line string) {
	// Ouverture du fichier
	file, _ := os.Open(banner + ".txt")
	// Création d'un scanner
	sc := bufio.NewScanner(file)
	lastLine := 0
	// Scan du fichier
	for sc.Scan() {
		// fmt.Println(sc.Text())
		lastLine++
		// Si la ligne correspond, récupération de la ligne
		if lastLine == lineNum {
			temp := sc.Text()
			file.Close()
			return temp
		}
	}
	return line
}

// fonction de print
func printmot(str string, banner string) {
	nbr := 0
	tmp := ""
	for i := 0; i < 8; i++ {
		for _, v := range str {
			// nbr permet de savoir a quel ligne se trouve la lettre
			nbr = i + (int(v)-32)*9 + 2
			// ajout de toutes les lignes dans tmp
			tmp = tmp + readLine(nbr, banner)
		}
		// print de la ligne
		fmt.Println(tmp)
		tmp = ""
	}
}

// fonction pour gérer les \n
func verificationRetourLigne(str string, banner string) {
	vide := true
	tmp := ""
	for i := 0; i < len(str); i++ {
		// si presence d'un '\n' et vide == true (ce n'ai pas le 1er '\n'), print d'un retour a la ligne
		if str[i] == '\\' && str[i+1] == 'n' && vide == true {
			fmt.Println()
			i++
			// si presence d'un '\n' et vide == false (1er '\n'), print du mot et passage de vide à true si suivi d'un autre '\n'
		} else if str[i] == '\\' && str[i+1] == 'n' {
			printmot(tmp, banner)
			time.Sleep(30 * time.Millisecond)
			tmp = ""
			i++
			vide = true
		} else {
			// print du mot et passage de vide à false pour gérer les '\n'
			tmp = tmp + string(str[i])
			vide = false
		}
	}
	if vide == false {
		printmot(tmp, banner)
		time.Sleep(30 * time.Millisecond)
	}

}
