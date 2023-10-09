package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func cal(str string) string {
	x, _ := strconv.Atoi(str)
	y, _ := strconv.Atoi(reverseString(str))

	return strconv.Itoa(x + y)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	str := sc.Text()

	for true {
		if str == reverseString(str) {
			fmt.Println(str)
			break
		}

		str = cal(str)
	}

}
