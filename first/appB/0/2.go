package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	str := strings.Split(sc.Text(), " ")

	card := make([]int, 0)

	for i := 0; i < n; i++ {
		number, _ := strconv.Atoi(str[i])
		card = append(card, number)
	}

	fmt.Println(getResult(&card))

}

func getResult(card *[]int) int {
	sort.Sort(sort.IntSlice(*card))

	sum := 0

	for i := 0; i < len(*card); i++ {
		if i == len(*card)-1 {
			sum += (*card)[i]
			break
		}

		if (*card)[i]+1 == (*card)[i+1] {
			continue
		}

		sum += (*card)[i]
	}

	return sum
}
