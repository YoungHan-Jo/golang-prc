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
	m, _ := strconv.Atoi(strs[1])

	sum := 0

	for i := 1; i < n; i++ {
		sc.Scan()
		if len, _ := strconv.Atoi(sc.Text()); len <= m {
			sum += len
		}
	}

	fmt.Println(sum)
}
