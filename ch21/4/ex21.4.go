package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

func main() {

	operator := getOperator("+")

	fmt.Println(operator(4, 5))

	result := getOperator("*")(3, 4)

	fmt.Println(result)

}
