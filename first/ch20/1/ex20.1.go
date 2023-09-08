package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("hello! I'm %s I'm %d years old", s.Name, s.Age)
}

func main() {

	student := Student{"Jack", 12}

	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())
}
