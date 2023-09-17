package main

import "fmt"

func main() {
	m := make(map[int]int)

	m[1] = 100
	m[2] = 200
	m[3] = 300

	delete(m, 3)

	for i := 0; i < 5; i++ {

		fmt.Printf("key : %v , value: ", i)

		if v, ok := m[i]; ok {
			fmt.Print(v)
		} else {
			fmt.Print("nil")
		}

		fmt.Println()
	}

}
