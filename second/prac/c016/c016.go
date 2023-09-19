package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	str := sc.Text()

	r := []rune(str)

	var result string = ""

	for _, v := range r {
		char := string(v)

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

	fmt.Println(result)
}
