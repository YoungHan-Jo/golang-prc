package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	point int
}

func (s *Student) Grade(answer, submission string) {
	if len(answer) != len(submission) {
		return
	}
	if answer == submission {
		s.point += 2
		return
	}

	answerRuneArr := []rune(answer)
	submitRuneArr := []rune(submission)

	count := 0
	for i := 0; i < len(answer); i++ {
		if answerRuneArr[i] != submitRuneArr[i] {
			count++
		}
	}
	if count > 1 {
		return
	}
	s.point += 1
}

func (s *Student) GetPoint() int {
	return s.point
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	student := Student{0}

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")

		answer := strs[0]
		submission := strs[1]
		student.Grade(answer, submission)
	}

	fmt.Println(student.GetPoint())
}
