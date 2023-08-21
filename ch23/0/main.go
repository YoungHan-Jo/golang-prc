package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	r := []rune(sc.Text())

	var result []rune

	for _, v := range r {
		if !strings.Contains("aeiouAEIOU", string(v)) {
			result = append(result, v)
		}
	}

	fmt.Println(string(result))

}
