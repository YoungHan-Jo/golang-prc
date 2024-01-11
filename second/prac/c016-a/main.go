package main

import (
	"bufio"
	"fmt"
	"os"
)

type Word struct {
	origin string
}

func (w *Word) changeLeet() string {

	runes := []rune(w.origin)

	var result string = ""

	for _, rune := range runes {
		char := string(rune)

		switch char {
		case "A":
			char = "4"
		case "E":
			char = "3"
		case "G":
			char = "6"
		case "I":
			char = "1"
		case "O":
			char = "0"
		case "S":
			char = "5"
		case "Z":
			char = "2"
		}

		result += char
	}

	return result
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	w := Word{sc.Text()}

	fmt.Println(w.changeLeet())
}
