package main

import "fmt"

func f() {
	fmt.Println("f() func start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recover -", r)
		}
	}()

	g()
	fmt.Println("f() func end")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("b can not b 0")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("program keeps running")
}
