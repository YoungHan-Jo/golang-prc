package main

import "fmt"

func main() {
	str := "hello 世界!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%T\t %d\t %c\n", str[i], str[i], str[i])
	}

	fmt.Println("-----------------")

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%T\t %d\t %c\n", runes[i], runes[i], runes[i])
	}

	fmt.Println("-----------------")

	for _, v := range str {
		fmt.Printf("%T\t %d\t %c\n", v, v, v)
	}
}
