package main

import (
	"fmt"
	"math"
)

const epsilon = 0.00001

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, c)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a, b))

	a = 0.000000000004
	b = 0.000000000002
	c = 0.000000000006

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a, b))
}

func equal(a, b float64) bool {
	if math.Abs(a-b) <= epsilon {
		return true
	}

	return false
}
