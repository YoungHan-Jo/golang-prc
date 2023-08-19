package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writerHello(writer Writer) {
	writer("Hello world")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	lf := func(msg string) {
		fmt.Fprintln(f, msg)
	}

	writerHello(lf)
}
