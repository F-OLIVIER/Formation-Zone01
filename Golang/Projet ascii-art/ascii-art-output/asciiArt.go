package asciiartoutput

import (
	"bufio"
	"os"
)

func Asciiart(input, banner string) [8]string {
	artFile, _ := os.Open(banner)
	tabstr := [8]string{}

	// lettersMap stock le modèle de chacune des lettres du texte
	lettersMap := make(map[byte][]string)
	scanner := bufio.NewScanner(artFile)
	letters := SortString(input)
	scannedLines := -1
	currentLetter := 0
	// construction de lettersMap.
	for scanner.Scan() && currentLetter < len(letters) {
		if scannedLines == (int(letters[currentLetter])-32)*9 {
			// la ligne scannée est la première de la prochaine lettre utilisée
			var model []string
			for i := 0; i < 8; i++ {
				model = append(model, scanner.Text())
				scanner.Scan()
				scannedLines++
			}
			lettersMap[letters[currentLetter]] = model
			currentLetter++
		}
		scannedLines++
	}
	for i := 0; i < 8; i++ {

		for _, letter := range input {
			tabstr[i] += lettersMap[byte(letter)][i]
		}

	}
	return tabstr
}
