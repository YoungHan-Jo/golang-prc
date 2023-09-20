package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var arr []string

	for i := 0; i < n; i++ {
		sc.Scan()

		arr = append(arr, sc.Text())
	}

	count := 0

	for i := 0; i < n-1; i++ {
		front := []rune(arr[i])
		back := []rune(arr[i+1])

		if n := len(arr[i]); front[n-1] != back[0] {
			fmt.Printf("%s %s", string(front[n-1]), string(back[0]))
			break
		}
		count++
	}

	if count == n-1 {
		fmt.Println("Yes")
	}
}
