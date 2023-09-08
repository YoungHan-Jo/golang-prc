package main

import "fmt"

func main() {
	var b float64 = 3.5

	x := int(b * 3)
	y := int(b) * 3

	fmt.Println(x, y)
}
