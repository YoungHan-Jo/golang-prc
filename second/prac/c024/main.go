package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	A int
	B int
}

func (c *Computer) Set(target *string, num *int) {
	switch *target {
	case "1":
		c.A = *num
		break
	case "2":
		c.B = *num
		break
	}
}

func (c *Computer) Add(num *int) {
	c.B = c.A + *num
}

func (c *Computer) Sub(num *int) {
	c.B = c.A - *num
}

func (c *Computer) Command(strs *[]string) {
	switch (*strs)[0] {
	case "SET":
		num, _ := strconv.Atoi((*strs)[2])
		c.Set(&(*strs)[1], &num)
	case "ADD":
		num, _ := strconv.Atoi((*strs)[1])
		c.Add(&num)
	case "SUB":
		num, _ := strconv.Atoi((*strs)[1])
		c.Sub(&num)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	computer := Computer{A: 0, B: 0}

	for i := 0; i < n; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")

		computer.Command(&strs)
	}

	fmt.Println(computer.A, computer.B)
}
