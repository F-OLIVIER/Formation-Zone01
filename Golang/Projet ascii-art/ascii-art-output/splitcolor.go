package asciiartoutput

import "strings"

func SplitColor(input string, colorLetter string) []string {
	inputSplit := make([]string, 0)

	for {
		index := strings.Index(input, colorLetter)
		newlineIndex := strings.Index(input, "\n")

		if index == -1 && newlineIndex == -1 {
			inputSplit = append(inputSplit, input)
			break
		}

		if index == -1 || (newlineIndex != -1 && newlineIndex < index) {
			if newlineIndex > 0 {
				inputSplit = append(inputSplit, input[:newlineIndex])
			}
			inputSplit = append(inputSplit, "")
			input = input[newlineIndex+1:]
			continue
		}

		if index > 0 {
			inputSplit = append(inputSplit, input[:index])
		}

		inputSplit = append(inputSplit, colorLetter)
		input = input[index+len(colorLetter):]
	}

	return inputSplit
}
