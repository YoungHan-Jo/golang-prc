package main

import "fmt"

func main() {
	temp := 30
	rain := 40

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("熱くて雨が降ります")
		} else if rain >= 20 {
			fmt.Println("熱くて湿気があります")
		} else {
			fmt.Println("お出かけしましょう")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("お出かけしないで")
	} else {
		fmt.Println("いい天気です")
	}
}
