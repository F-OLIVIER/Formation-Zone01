package asciiartcolor

import (
	// "os"
	"fmt"
	"io/ioutil"
	"strings"
)

func AddToTable(v byte, strTable [8]string, data []string) [8]string {
	count := 0

	for count < 8 {
		/*
			Each character is 8 lines.
			We loop 8 times to retrieve the parts of the character.
			The formula to retrieve the part of the character in the file is:

			value ascii of the character - 32 (the first printable character is space at 32) * 9 + counter + 1.

			We add what we retrieve in the string table in each string.
			We return the string table once we're done.
		*/
		temp := []rune{}
		for _, v := range strTable[count] {
			temp = append(temp, v)
		}
		for _, v := range data[(int(v)-32)*9+count+1] {
			if v != 13 {
				temp = append(temp, v)
			}
		}
		strTemp := ""
		for _, v := range temp {
			strTemp += string(v)
		}
		strFinal := ""
		for _, v := range strTemp {
			if v == ' ' {
				strFinal += " "
			} else {
				strFinal += string(v)
			}
		}
		strTable[count] = strFinal
		count++
	}

	return strTable
}

/* func printStr(strTable [8]string) {
	for _, v := range strTable {
		fmt.Println(v)
	}
} */

func Asciiart(word, radio string) [8]string {
	/* if len(os.Args) != 2 {
		os.Exit(1)
	} */
	byte, _ := ioutil.ReadFile(radio)
	data := strings.Split(string(byte), "\n")
	strTable := [8]string{} // We prepare an empty table of string that will be used to store the strings to print
	// word := os.Args[1]
	for i := 0; i < len(word); i++ { // We loop for each character in the argument
		v := word[i]
		if v == '\\' && len(word)-1 != i { // If we find a \ & the index is not on the last character
			if word[i+1] == 'n' { // If the next character is an n
				if len(strTable[0]) > 0 {
					return strTable // If we find a \n and the table of string has something ot be printed
					// printStr(strTable)
				} else {
					fmt.Println() // Otherwirse we just print an empty line
				}
				strTable = [8]string{} // We reset the table of string
				i++
			} else {
				strTable = AddToTable(v, strTable, data) // We call the function addToTable to add each part of a word in the table of string
			}
		} else {
			strTable = AddToTable(v, strTable, data) // We call the function addToTable to add each part of a word in the table of string
		}
	}
	if len(strTable[0]) > 0 {
		return strTable
		// printStr(strTable) //if we're at the end of the word and there's something to print we print it
	}
	return strTable
}
