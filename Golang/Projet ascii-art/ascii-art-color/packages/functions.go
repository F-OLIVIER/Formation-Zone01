package functions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func First(b int, _ error) int {
	return b
}

// Fonction qui génére un tableau de l'ascii
func GetAsciiChar(asciiSample []string, charIndex int) []string {

	result := []string{}
	start := (charIndex * 8) + charIndex

	for i := start; i <= start+8; i++ {
		result = append(result, asciiSample[i])
	}

	return result
}

func HexToColor(x string) string {

	if !strings.HasPrefix(x, "#") {
		if !strings.HasPrefix(x, "hex(") || !strings.HasSuffix(x, ")") {
			return "invalidHEX"
		}
	}

	x = strings.Replace(strings.Replace(x, ")", "", 1), "hex(", "", 1)

	var intArr []int // gestion du string d'entrée en int

	for index, i := range x {

		switch true {
		case i == '#' && index == 0:
			{
				continue
			}
		case i >= 'a' && i <= 'f':
			{
				intArr = append(intArr, int(i-'a'+10))
			}
		case i >= 'A' && i <= 'F':
			{
				intArr = append(intArr, int(i-'A'+10))
			}
		case i >= '0' && i <= '9':
			{
				intArr = append(intArr, int(i-'0'))
			}
		default:
			{
				return "invalidHEX"
			}
		}

	}

	if len(intArr) != 6 {
		return "invalidHEX"
	}

	// Calcul de convertion
	// r = Première valeur x 16 + deuxième valeur = niveau de rouge de 0 à 255
	// g = Troisième valeur x 16 + quatrième valeur = niveau de vert de 0 à 255
	// b = Cinquième valeur x 16 + sixième valeur = niveau de bleu de 0 à 255
	r := 16*intArr[0] + intArr[1]
	g := 16*intArr[2] + intArr[3]
	b := 16*intArr[4] + intArr[5]

	// on génére le string a inclure pour géré la couleur
	return "\x1b[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

func RgbToColor(str string) string {

	if !strings.HasPrefix(str, "rgb(") || !strings.HasSuffix(str, ")") {
		return "invalidRGB"
	}

	rgb := strings.Split(strings.Replace(strings.Replace(str, ")", "", 1), "rgb(", "", 1), ",")

	if len(rgb) != 3 {
		return "invalidRGB"
	}

	for _, item := range rgb {
		resAtoi, err := strconv.Atoi(item)
		if err != nil || resAtoi > 255 || resAtoi < 0 {
			return "invalidRGB"
		}
	}

	return "\x1b[38;2;" + rgb[0] + ";" + rgb[1] + ";" + rgb[2] + "m"

}

func HslToColor(str string) string {

	if !strings.Contains(str, "%") {
		if !strings.HasPrefix(str, "hsl(") || !strings.HasSuffix(str, ")") {
			return "invalidHSL"
		}
	}

	hsl := strings.Split(strings.Replace(strings.Replace(strings.Replace(str, ")", "", 1), "hsl(", "", 1), "%", "", 2), ",")

	if len(hsl) != 3 {
		return "invalidHSL"
	}

	hslFloatArr := []float64{}

	for index, item := range hsl {

		f, err := strconv.ParseFloat(item, 32)
		if err != nil || ((index == 0 && f > 360) || ((index == 1 || index == 2) && f > 100)) || f < 0 {
			return "invalidHSL"
		}
		hslFloatArr = append(hslFloatArr, f)
	}

	h := hslFloatArr[0]
	s := hslFloatArr[1]
	l := hslFloatArr[2]

	if h == 0 && s == 0 && l == 0 {
		return "\x1b[38;2;0;0;0m"
	}

	if hslFloatArr[1] == 0 {
		return "\x1b[38;2;" + FloatToString(l) + ";" + FloatToString(l) + ";" + FloatToString(l) + "m"
	}

	s /= 100
	l /= 100

	a := float64(s) * math.Min(float64(l), float64(1-l))

	r := int(255 * F(0, l, a, h))
	g := int(255 * F(8, l, a, h))
	b := int(255 * F(4, l, a, h))

	return "\x1b[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"

}

func FloatToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

func K(n float64, h float64) float64 {
	return math.Mod(n+h/30, 12)
}

func F(n float64, l float64, a float64, h float64) float64 {
	return l - a*math.Max(-1, math.Min(K(n, h)-3, math.Min(9-K(n, h), 1)))
}

func CheckParamAndReturnColor(str string) string {

	if !strings.HasPrefix(str, "--color=") {
		return "invalidFormat"
	}

	color := strings.TrimPrefix(str, "--color=")

	if color == "" {
		return "invalidFormat"
	}

	availableColorNames := [][]string{{"red", "\033[31m"}, {"green", "\033[32m"}, {"blue", "\033[34m"}, {"yellow", "\033[33m"}, {"purple", "\033[35m"}, {"cyan", "\033[36m"}, {"gray", "\033[37m"}, {"orange", "\033[38;2;255;165;0m"}}

	hexRes := HexToColor(color)
	rgbRes := RgbToColor(color)
	hslRes := HslToColor(color)

	// Vérification de la validité de la couleur par rapport au couleur crée

	switch true {
	case hexRes == "invalidHEX" && rgbRes == "invalidRGB" && hslRes == "invalidHSL":
		{
			for i, item := range availableColorNames {
				if item[0] == color {
					return item[1]
				} else if i == (len(availableColorNames) - 1) {
					return "invalidColor"
				}
			}
		}
	case hexRes != "invalidHEX":
		{
			return hexRes
		}
	case hslRes != "invalidHSL":
		{
			return hslRes
		}
	case rgbRes != "invalidRGB":
		{
			return rgbRes
		}
	}

	return "invalidColor"

}

func FormatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
}

func OtherColorsFilterContains(arr [][]string, char string) bool {

	for _, item := range arr {
		if strings.Contains(item[1], char) {
			return true
		}
	}

	return false

}
