package Check

import (
	"fmt"
	Struct "my-ls-1/package/struct"
	"os"
)

func CheckEntrer(args []string) []Struct.Option {
	var listeOption []Struct.Option
	var option Struct.Option
	// countTirer := 0
	option.OptionPetitA = false
	option.OptionPetitL = false
	option.OptionGrandR = false
	option.OptionPetitR = false
	option.OptionPetitT = false
	for _, arg := range args {
		if arg[0] == '-' && len(arg) == 1 {
			option.Argument = append(option.Argument, arg)
		} else if arg[0] == '-' {
			// countTirer++
			// if countTirer > 1 {
			// 	listeOption = append(listeOption, option)
			// 	option.OptionPetitA = false
			// 	option.OptionPetitL = false
			// 	option.OptionGrandR = false
			// 	option.OptionPetitR = false
			// 	option.OptionPetitT = false
			// 	option.Argument = []string{}
			// }
			for _, letter := range arg {
				if letter == '-' {
					// ne rien faire (si absent, le else pose problème pour le 1er '-')
				} else if letter == 'a' {
					option.OptionPetitA = true
				} else if letter == 'l' {
					option.OptionPetitL = true
				} else if letter == 'R' {
					option.OptionGrandR = true
				} else if letter == 'r' {
					option.OptionPetitR = true
				} else if letter == 't' {
					option.OptionPetitT = true
				} else {
					fmt.Println("ls : option invalide -- '" + string(letter) + "'\nSaisissez « ls --help » pour plus d'informations.")
					os.Exit(0)
				}
			}
		} else {
			option.Argument = append(option.Argument, arg)
		}
	}
	// ajout de la derniére struct d'option
	listeOption = append(listeOption, option)
	return listeOption
}
