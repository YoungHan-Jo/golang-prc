package main

import "fmt"

type number int

func main() {
	var i number = 0

	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}
