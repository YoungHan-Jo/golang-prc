package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs[0])
	x, _ := strconv.Atoi(strs[1])
	y, _ := strconv.Atoi(strs[2])

	var result []string

	for i := 1; i <= n; i++ {
		sum := 0

		if i%x == 0 {
			sum += 1
		}

		if i%y == 0 {
			sum += 2
		}

		switch sum {
		case 0:
			result = append(result, "N")
		case 1:
			result = append(result, "A")
		case 2:
			result = append(result, "B")
		case 3:
			result = append(result, "AB")
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
