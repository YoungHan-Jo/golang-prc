package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type User struct {
	num  int
	name string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var users []User

	for i := 0; i < n; i++ {
		sc.Scan()

		name := sc.Text()

		re := regexp.MustCompile("[0-9]+")
		num, _ := strconv.Atoi(re.FindAllString(name, -1)[0])

		users = append(users, User{num, name})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].num < users[j].num
	})

	for _, u := range users {
		fmt.Println(u.name)
	}

}
