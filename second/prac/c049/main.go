package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Elevator struct {
	floor int
}

func (e *Elevator) moveTo(nextFloor *int) int {
	move := int(math.Abs(float64(*nextFloor - e.floor)))
	e.floor = *nextFloor
	return move
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	e := Elevator{floor: 1}

	sum := 0

	for i := 0; i < n; i++ {
		sc.Scan()
		f, _ := strconv.Atoi(sc.Text())
		sum += e.moveTo(&f)
	}

	fmt.Println(sum)
}
