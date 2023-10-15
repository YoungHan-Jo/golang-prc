package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator(point, day, count int) int {
	sum := point * count
	if day >= 10 {
		return 0
	}
	if day > 1 {
		return int(float64(sum) * 0.8)
	}
	return sum
}

func grade(points int) string {
	if points >= 80 {
		return "A"
	}
	if points >= 70 {
		return "B"
	}
	if points >= 60 {
		return "C"
	}
	return "D"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(strs[0])
	n, _ := strconv.Atoi(strs[1])
	point := 100 / n

	for i := 0; i < k; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		day, _ := strconv.Atoi(strs[0])
		count, _ := strconv.Atoi(strs[1])

		fmt.Println(grade(calculator(point, day, count)))
	}
}
