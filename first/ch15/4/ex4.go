package main

import "fmt"

func main() {
	var char rune = '日'

	fmt.Printf("%T\n", char) // int32
	fmt.Println(char)
	fmt.Printf("%c\n", char)

	str1 := "あいうえお" // japanese 3byte
	str2 := "abcde" // english 1byte

	fmt.Println(len(str1)) // 15   // size of memory
	fmt.Println(len(str2)) // 5

	runes := []rune(str1) // string to []rune
	fmt.Println(len(runes))
}
