package main

import "fmt"

func main() {
	c, isSuccess := Divide(4, 0)
	fmt.Println(c, isSuccess)
	d, isSuccess := Divide(4, 2)
	fmt.Println(d, isSuccess)
}

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
