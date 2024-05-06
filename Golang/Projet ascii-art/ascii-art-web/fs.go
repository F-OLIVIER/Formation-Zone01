package fs

import (
	"io/ioutil"
	"strings"
)

func PrintAll(word string, option string) string {
	var res string
	res += "\n"

	byte, _ := ioutil.ReadFile("./templates/" + option + ".txt")

	byte2 := strings.ReplaceAll(string(byte), "\r", "")

	data := strings.Split(byte2, "\n")

	if strings.Contains(word, "\\n") {
		words := strings.Split(word, "\\n")

		for i := 0; i < len(words); i++ {
			switch {
			case i != len(words)-1 && words[i] != "":
				res += GenAsciiArt(words[i], data)
			case i != len(words)-1 && words[i] == "":
				res += "\n"
			case i == len(words)-1:
				res += GenAsciiArt(words[i], data)
			}
		}
	} else {
		res += GenAsciiArt(word, data)
	}
	return res
}

func GenAsciiArt(word string, data []string) (res string) {
	var strTable [8]string
	for i := 0; i < len(word); i++ {
		count := 0
		for count < 8 {
			strTable[count] += data[(int(word[i])-32)*9+count+1]
			count++
		}
	}
	for _, v := range strTable {
		res += v + "\n"
	}
	return res
}
