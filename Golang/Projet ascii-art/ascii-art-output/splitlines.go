package asciiartoutput

func SplitLines(input string) []string {
	// Sépare l'entrée en un tableau de string en fonction des sauts de ligne avec \n
	var res []string
	i := 0
	res = append(res, "")
	antiSlash := false
	for j := 0; j < len(input); j++ {
		if input[j] == 'n' {
			if antiSlash {
				res[i] = res[i][:len(res[i])-1]
				i++
				res = append(res, "")
			} else {
				res[i] += "n"
			}
		} else if input[j] == '\n' {
			i++
			res = append(res, "")
		} else {
			res[i] += string(input[j])
			if input[j] == '\\' {
				antiSlash = true
			} else {
				antiSlash = false
			}
		}
	}
	return res
}
