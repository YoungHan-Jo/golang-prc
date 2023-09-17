package main

import (
	"bufio"
	"container/list"
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

	var boxMaxCounts []int

	groupList := list.New()

	var result []int

	all := 0

	for i := 0; i < n; i++ {
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		boxMaxCounts = append(boxMaxCounts, c)
		result = append(result, 0)
	}

	for i := 0; i < m; i++ {
		sc.Scan()

		c, _ := strconv.Atoi(sc.Text())
		groupList.PushBack(c)

		all += c
	}

	element := groupList.Front()

	left := false
	count := 0

	for all != 0 {
		for i := 0; i < n; i++ {
			if element == nil {
				break
			}

			if !left {
				count = element.Value.(int)
			}

			boxMax := boxMaxCounts[i]

			if boxMax >= count {
				all -= count
				result[i] += count
				element = element.Next()
				left = false
			} else {
				count -= boxMax
				result[i] += boxMax
				all -= boxMax
				left = true
			}
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}

}
