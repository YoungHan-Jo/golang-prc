package main

import "fmt"

func find45(a, b int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}

	return 0, false
}

func main() {
	a := 0
	b := 0

	for ; a <= 9; a++ {
		var isFound bool
		if b, isFound = find45(a, b); isFound {
			break
		}
	}

	fmt.Println(a, b)
}
