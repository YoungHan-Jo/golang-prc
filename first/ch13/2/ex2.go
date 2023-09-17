package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

type VIPUser2 struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"UserA", "AAA", 10}
	vip := VIPUser{
		user,
		1,
		1000,
	}
	fmt.Println("vip:", vip)
	fmt.Println(vip.UserInfo.Name)

	vip2 := VIPUser2{
		User{"UserB", "BBB", 20},
		2,
		2000,
	}

	fmt.Println(vip2.Name)
}
