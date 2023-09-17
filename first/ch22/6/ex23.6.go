package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"Pen", 200}
	m[28] = Product{"Eraser", 100}
	m[44] = Product{"SharpPencil", 300}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
