package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house1 House
	house1.Address = "Tokyo"
	house1.Size = 72
	house1.Price = 5.6
	house1.Type = "mention"

	fmt.Println("house1:", house1)

	var house2 House = House{"Seoul", 108, 10.2, "apartment"}
	fmt.Println("house2:", house2)

	var house3 House = House{
		"Sydney",
		99,
		9.1,
		"house",
	}
	fmt.Println("house3:", house3)

}
