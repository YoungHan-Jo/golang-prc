package main

import "fmt"

func main() {

	temp := 10

	switch {
	case temp > 20:
		fmt.Println("over 20")
	case temp > 10:
		fmt.Println("over 10")
	case temp > 0:
		fmt.Println("over 0")
	}
}
