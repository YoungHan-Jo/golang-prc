package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b can not be 0")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
