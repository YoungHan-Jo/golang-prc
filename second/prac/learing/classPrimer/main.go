package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee interface {
	GetNum() string
	GetName() string
	ChangeName(*string)
	ChangeNum(*string)
}

func NewEmployee(number, name string) Employee {
	return &employee{number, name}
}

type employee struct {
	number string
	name   string
}

func (e *employee) GetNum() string {
	return e.number
}

func (e *employee) GetName() string {
	return e.name
}

func (e *employee) ChangeName(newName *string) {
	e.name = *newName
}

func (e *employee) ChangeNum(newNum *string) {
	e.number = *newNum
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var employees []Employee

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")

		switch strs[0] {
		case "make":
			newEmp := NewEmployee(strs[1], strs[2])
			employees = append(employees, newEmp)
		case "getnum":
			index, _ := strconv.Atoi(strs[1])
			fmt.Println(employees[index-1].GetNum())
		case "getname":
			index, _ := strconv.Atoi(strs[1])
			fmt.Println(employees[index-1].GetName())
		case "change_num":
			index, _ := strconv.Atoi(strs[1])
			employees[index-1].ChangeNum(&strs[2])
		case "change_name":
			index, _ := strconv.Atoi(strs[1])
			employees[index-1].ChangeName(&strs[2])
		}
	}
}
