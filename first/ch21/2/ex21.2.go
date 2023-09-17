package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("1")
	defer f.Close()
	defer fmt.Println("2")

	fmt.Println(`Write "Hello World" into file`)
	fmt.Fprintln(f, "Hello World")
}
