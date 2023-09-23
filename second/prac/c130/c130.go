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

	n, _ := strconv.Atoi(sc.Text())

	var result []int

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		first := strs[0]
		if second := strs[1]; second == first && second == "y" {
			continue
		}
		result = append(result, i+1)
	}

	fmt.Println(len(result))
	for _, v := range result {
		fmt.Println(v)
	}
}
