package push_swap

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func VerifyAndConvert() (intArr, intArr1 []int) {
	// vérifie la quantité d'arguments
	if len(os.Args) < 2 || len(os.Args) > 2 {
		os.Exit(0)
	}
	var args string = os.Args[1]
	// On split notre string
	strArr := strings.Split(args, " ")
	// on effectue la convertion en int
	for _, s := range strArr {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("ERROR : Caractére autres qu'un int")
			os.Exit(1)
		}
		intArr = append(intArr, num)

	}
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr); j++ {
			if intArr[i] == intArr[j] && i != j {
				fmt.Println("ERROR : int identique")
				os.Exit(1)
			}
		}
	}

	// On verifie si le tableau est deja dans l'ordre
	for i := 0; i < len(intArr)-1; i++ {
		if intArr[i] > intArr[i+1] {
			return intArr, intArr1
		}
	}
	os.Exit(0)
	return intArr, intArr1
}

func Pipe() (listOperation []string) {
	// lecture du contenu du pipe
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" && scanner.Text() != "-e" {
			if strings.Contains(scanner.Text(), "Nombre d'opération") {
				break
			}
			tmp := strings.ReplaceAll(strings.ToUpper(scanner.Text()), "-E ", "")
			listOperation = append(listOperation, tmp)
		}

	}

	// gestion d'un pipe vide
	if len(listOperation) == 0 {
		fmt.Println("KO")
		os.Exit(0)
	}

	return listOperation
}
