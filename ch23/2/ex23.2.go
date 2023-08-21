package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("float can not be negative. f: %f", f)
		return 0, errors.New("float can not be negative.")
	}

	return math.Sqrt(f), nil
}

func main() {
	f, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%f", f)

}
