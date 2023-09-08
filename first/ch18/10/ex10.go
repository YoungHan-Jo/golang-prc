package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println(slice)

	slice1 := []int{1, 2, 3, 4, 5, 6}

	slice1 = append(slice1[:idx], append([]int{100}, slice1[idx:]...)...)
	fmt.Println(slice1)
}
