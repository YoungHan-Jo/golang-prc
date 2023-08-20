package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for i := 0; i < r.Len()*2; i++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}

}
