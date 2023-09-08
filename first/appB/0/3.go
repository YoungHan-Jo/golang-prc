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

	strs1 := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs1[0])
	month, _ := strconv.Atoi(strs1[1])

	sc.Scan()
	strs2 := strings.Split(sc.Text(), " ")
	init, _ := strconv.Atoi(strs2[0])
	person, _ := strconv.Atoi(strs2[1])
	price, _ := strconv.Atoi(strs2[2])

	sum := 0

	for i := 0; i < n; i++ {
		sc.Scan()
		if amount, _ := strconv.Atoi(sc.Text()); init+person*month > amount*price {
			sum++
		}
	}

	fmt.Println(sum)
}
