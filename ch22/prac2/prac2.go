package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	m := make(map[int]Student)
	m[38] = Student{"Annie", 78}
	m[45] = Student{"Bell", 56}

	fmt.Println(m[38])
}
