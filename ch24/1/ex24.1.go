package main

import (
	"fmt"
	"time"
)

func PrintAlphabet() {
	alphabet := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	for _, v := range alphabet {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintAlphabet()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
	fmt.Println("end")
}
