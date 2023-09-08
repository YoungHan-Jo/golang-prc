package main

import "fmt"

func main() {

	switch day := "thursday"; day {
	case "monday", "thursday":
		fmt.Println("study")
	case "wednesday", "friday":
		fmt.Println("exercise")
	}
}
