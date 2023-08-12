package main

import "fmt"

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)

func main() {
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)

	fmt.Println("=======")
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)
}
