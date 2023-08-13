package main

import "fmt"

func main() {

	var array = [...]int{1, 2, 3}
	var slice = []int{1, 2, 3}
	fmt.Println("array :", array)
	fmt.Println("slice :", slice)

	var slice2 = make([]int, 3)
	fmt.Println("slice using make :", slice2)

	fmt.Println("-------range-------")
	var slice3 = []int{1, 2, 3}

	for i, v := range slice3 {
		slice3[i] *= 2
		fmt.Println(v)
	}

	for _, v := range slice3 {
		fmt.Println(v)
	}
}
