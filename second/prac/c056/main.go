package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	point  int
	absent int
}

func (s *Student) isPassed(deadline int) bool {
	grade := s.point - s.absent*5

	if grade < 0 {
		grade = 0
	}
	return grade >= deadline
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs[0])
	deadline, _ := strconv.Atoi(strs[1])

	for i := 1; i < n+1; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")

		point, _ := strconv.Atoi(strs[0])
		absent, _ := strconv.Atoi(strs[1])

		student := Student{point, absent}

		if student.isPassed(deadline) {
			fmt.Println(i)
		}
	}

}
