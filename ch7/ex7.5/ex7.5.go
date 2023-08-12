package main

import "fmt"

func Divide(a, b int) (result int, isSuccess bool) {
	if b == 0 {
		result = 0
		isSuccess = false
		return
	}
	result = a / b
	isSuccess = true
	return
}

func main() {
	c, isSuccess := Divide(4, 0)
	fmt.Println(c, isSuccess)
	d, isSuccess := Divide(4, 2)
	fmt.Println(d, isSuccess)
}
