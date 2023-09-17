package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(5))
	var b = 20
	fmt.Println(myInt(b).add(5))
}
