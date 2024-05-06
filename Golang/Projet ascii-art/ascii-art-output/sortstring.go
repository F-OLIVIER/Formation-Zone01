package asciiartoutput

import (
	"sort"
	"strings"
)

func SortString(text string) string {
	// décompose un string en la suite ordonnée des letters qui la composent
	// ex : "caab" -> "abc"
	res := ""
	for _, letter := range text {
		if letter != '\n' && !strings.Contains(res, string(letter)) {
			res += string(letter)
		}
	}
	s := strings.Split(res, "")
	sort.Strings(s)
	return strings.Join(s, "")

}
