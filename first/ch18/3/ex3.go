package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 15; i++ {
		slice = append(slice, i)
	}

	fmt.Println("slice :", slice)
}
