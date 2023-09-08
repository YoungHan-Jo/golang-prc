package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type opFunc func(int, int) int // alias

func getOperator(op string) opFunc {
	switch op {
	case "+":
		return add
	case "*":
		return mul
	default:
		return nil
	}
}

func main() {
	var operator opFunc
	operator = getOperator("*")

	var result = operator(2, 5)
	fmt.Println(result)
}
