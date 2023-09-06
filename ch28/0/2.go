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

	result := 0

	for i := 0; i < n; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")

		major := strs[0]
		eng, _ := strconv.Atoi(strs[1])
		math, _ := strconv.Atoi(strs[2])
		science, _ := strconv.Atoi(strs[3])
		language, _ := strconv.Atoi(strs[4])
		history, _ := strconv.Atoi(strs[5])

		if eng+math+science+language+history < 350 {
			continue
		}

		if major == "l" {
			if language+history >= 160 {
				result++
			}
		}

		if major == "s" {
			if math+science >= 160 {
				result++
			}
		}
	}

	fmt.Println(result)
}
