package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func Atoi(input string) (int, error) {

	input = strings.Replace(input, " ", "", -1)

	v, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
