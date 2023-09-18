package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func strToInteger(str *string) int {
	n := strings.Count(*str, "<")
	m := strings.Count(*str, "/")
	return n*10 + m
}

func result(strs *[]string) int {
	sum := 0
	for _, v := range *strs {
		sum += strToInteger(&v)
	}

	return sum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), "+")

	fmt.Println(result(&strs))
}
