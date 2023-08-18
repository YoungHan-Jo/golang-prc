package main

import "fmt"

func sum(nums ...int) int {

	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum())
}
