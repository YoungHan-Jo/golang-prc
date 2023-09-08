package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	cal := strs[1]

	var position int

	a, err := strconv.Atoi(strs[0])
	if err != nil {
		position = 1
	}

	b, err := strconv.Atoi(strs[2])
	if err != nil {
		position = 2
	}

	c, err := strconv.Atoi(strs[4])
	if err != nil {
		position = 3
	}

	switch position {
	case 1:
		if cal == "+" {
			fmt.Println(c - b)
		} else {
			fmt.Println(c + b)
		}
	case 2:
		if cal == "+" {
			fmt.Println(c - a)
		} else {
			fmt.Println(a - c)
		}
	case 3:
		if cal == "+" {
			fmt.Println(a + b)
		} else {
			fmt.Println(a - b)
		}
	}
}
