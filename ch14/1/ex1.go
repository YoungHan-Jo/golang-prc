package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Println(p)
	fmt.Println(*p)

	*p = 300

	fmt.Println(a)
}
