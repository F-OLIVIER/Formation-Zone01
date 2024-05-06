package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	argument := os.Args[1:]
	if len(argument) != 1 {
		fmt.Println("Only one argument needed, please")
		return
	}
	str := argument[0]
	verificationRetourLigne(str)
}

// fonction pour lire une ligne d'un fichier
func readLine(lineNum int) (line string) {
	// Ouverture du fichier
	file, _ := os.Open("standard.txt")
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
func printmot(str string) {
	nbr := 0
	tmp := ""
	for i := 0; i < 8; i++ {
		for _, v := range str {
			// nbr permet de savoir a quel ligne se trouve la lettre
			nbr = i + (int(v)-32)*9 + 2
			// ajout de toutes les lignes dans tmp
			tmp = tmp + readLine(nbr)
		}
		// print de la ligne
		fmt.Println(tmp)
		tmp = ""
	}
}

// fonction pour gérer les \n
func verificationRetourLigne(str string) {
	vide := true
	tmp := ""
	for i := 0; i < len(str); i++ {
		// si presence d'un '\n' et vide == true (ce n'ai pas le 1er '\n'), print d'un retour a la ligne
		if str[i] == '\\' && str[i+1] == 'n' && vide == true {
			fmt.Println()
			i++
			// si presence d'un '\n' et vide == false (1er '\n'), print du mot et passage de vide à true si suivi d'un autre '\n'
		} else if str[i] == '\\' && str[i+1] == 'n' {
			printmot(tmp)
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
		printmot(tmp)
		time.Sleep(30 * time.Millisecond)
	}

}
