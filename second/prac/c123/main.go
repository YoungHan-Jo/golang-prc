package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	age       int
	beanCount int
}

func (p *Person) getBeans(amount *int) {
	p.beanCount += *amount

	if p.beanCount > p.age {
		p.beanCount = p.age
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var people []Person

	for i := 0; i < n; i++ {
		sc.Scan()
		age, _ := strconv.Atoi(sc.Text())
		people = append(people, Person{age: age, beanCount: 0})
	}

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	for i := 0; i < m; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")

		start, _ := strconv.Atoi(strs[0])
		end, _ := strconv.Atoi(strs[1])
		amount, _ := strconv.Atoi(strs[2])

		for i := start - 1; i < end; i++ {
			people[i].getBeans(&amount)
		}
	}

	for _, v := range people {
		fmt.Println(v.beanCount)
	}

}
