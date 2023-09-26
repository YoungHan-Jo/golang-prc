package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Band struct {
	name      string
	schedules []int
}

func (b *Band) addSchedule(day *int) {
	b.schedules = append(b.schedules, *day)
}

func (b *Band) memoCalendar(cal *[31]string) {
	isChange := false
	for _, v := range b.schedules {
		day := v - 1
		if cal[day] == "x" {
			cal[day] = b.name
			continue
		}
		if isChange {
			cal[day] = b.name
		}
		isChange = !isChange
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	bandA := Band{name: "A"}
	bandB := Band{name: "B"}

	for i := 0; i < n; i++ {
		sc.Scan()
		day, _ := strconv.Atoi(sc.Text())
		bandA.addSchedule(&day)
	}

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	for i := 0; i < m; i++ {
		sc.Scan()
		day, _ := strconv.Atoi(sc.Text())
		bandB.addSchedule(&day)
	}

	var result [31]string

	for i := 0; i < len(result); i++ {
		result[i] = "x"
	}

	bandA.memoCalendar(&result)
	bandB.memoCalendar(&result)

	for _, v := range result {
		fmt.Println(v)
	}
}
