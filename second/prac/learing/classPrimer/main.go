package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name  string
	old   string
	birth string
	state string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var students []student

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		students = append(students, student{strs[0], strs[1], strs[2], strs[3]})
	}

	sc.Scan()
	old := sc.Text()

	for _, v := range students {
		if v.old == old {
			fmt.Println(v.name)
			break
		}
	}
}
