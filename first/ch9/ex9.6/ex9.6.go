package main

import "fmt"

func getMyAge() (int, bool) {
	return 30, true
}

func main() {
	if age, isOk := getMyAge(); isOk && age < 20 {
		fmt.Println("you are young", age)
	} else if isOk && age < 30 {
		fmt.Println("you are nice age", age)
	} else if isOk {
		fmt.Println("you are beautiful age", age)
	} else {
		fmt.Println("error")
	}
}
