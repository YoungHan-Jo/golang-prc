package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pullNumber(inputNumber int) int {
	if inputNumber == 0 {
		return 24
	}

	return (inputNumber + 2) * 2
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	str := strings.ReplaceAll(sc.Text(), "-", "")
	runes := []rune(str)

	sum := 0

	for i := 0; i < len(runes); i++ {
		n, _ := strconv.Atoi(string(runes[i]))
		sum += pullNumber(n)
	}

	fmt.Println(sum)
}
