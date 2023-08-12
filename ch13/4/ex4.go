package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	// user := User{"UserA", "AAA", 10, 3}
	vip := VIPUser{
		User{"UserB", "BBB", 20, 8},
		2000,
		5,
	}
	fmt.Println(vip.Level)
	fmt.Println(vip.User.Level)
}
