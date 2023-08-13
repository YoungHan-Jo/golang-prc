package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Println(p1 == p2)
	fmt.Println(p2 == p3)

	var p *int

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("nil")
	}
}
