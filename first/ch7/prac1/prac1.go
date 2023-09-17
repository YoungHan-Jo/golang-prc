package main

import "fmt"

func Multiple(a, b int) int {
	return a * b
}

func main() {
	var a int = 2
	var b int = 4

	mul := Multiple(a, b)

	fmt.Println(mul)
}
