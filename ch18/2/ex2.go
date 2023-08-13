package main

import "fmt"

func main() {

	var slice = []int{1, 2, 3}

	slice2 := append(slice, 4, 5, 6, 7)
	slice = append(slice, 4, 5, 6, 7)

	fmt.Println(slice)
	fmt.Println(slice2)
}
