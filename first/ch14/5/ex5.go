package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var user = User{name, age}
	return &user
}

func main() {
	a := NewUser("AAA", 20)
	fmt.Println(a)
}
