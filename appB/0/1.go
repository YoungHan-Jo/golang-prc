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

	str := strings.Split(sc.Text(), " ")

	card := make([]int, 0)

	for _, v := range str {
		number, _ := strconv.Atoi(v)
		card = append(card, number)
	}

	fmt.Println(getResult(&card))
}

func getResult(card *[]int) int {
	sort.Sort(sort.IntSlice(*card))
	return (*card)[0] + (*card)[1] + (*card)[2]*10 + (*card)[3]*10
}
