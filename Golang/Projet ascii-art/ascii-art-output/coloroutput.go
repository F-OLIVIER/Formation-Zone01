package asciiartoutput

func ColorOutput(color string) string {

	if color == "red" || color == "Red" || color == "RED" || color == "rgb(255, 0, 0)" {
		return Red
	} else if color == "green" || color == "Green" || color == "GREEN" || color == "rgb(0, 255, 0)" {
		return Green
	} else if color == "blue" || color == "Blue" || color == "BLUE" || color == "rgb(0, 0, 255)" {
		return Blue
	} else if color == "yellow" || color == "Yellow" || color == "YELLOW" || color == "rgb(255, 255, 0)" {
		return Yellow
	} else if color == "purple" || color == "Purple" || color == "PURPLE" || color == "rgb(255, 0, 255)" {
		return Purple
	} else if color == "cyan" || color == "Cyan" || color == "CYAN" || color == "rgb(0, 255, 255)" {
		return Cyan
	} else if color == "gray" || color == "Gray" || color == "GRAY" || color == "rgb(0, 0, 0)" {
		return Gray
	}

	return Reset
}
