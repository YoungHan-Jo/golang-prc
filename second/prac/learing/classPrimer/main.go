package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	nickname string
	old      string
	birth    string
	state    string
}

func (u *User) changeName(name *string) {
	u.nickname = *name
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs[0])
	k, _ := strconv.Atoi(strs[1])

	var users []User

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		users = append(users, User{strs[0], strs[1], strs[2], strs[3]})
	}

	for i := 0; i < k; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		index, _ := strconv.Atoi(strs[0])

		users[index-1].changeName(&strs[1])
	}

	for _, v := range users {
		fmt.Printf("%s %s %s %s\n", v.nickname, v.old, v.birth, v.state)
	}
}
