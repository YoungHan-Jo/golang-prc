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

	start, _ := strconv.Atoi(strs[0])
	gap, _ := strconv.Atoi(strs[1])

	for i := 0; i < 10; i++ {
		fmt.Printf("%d", start+gap*i)
		if i != 9 {
			fmt.Print(" ")
		}
	}
}
