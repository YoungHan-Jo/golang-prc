package main

import (
	"time"
)

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	parcel := &Parcel{}
	parcel.Pdt = pdt
	parcel.ShippedTime = time.Now()
	return parcel
}

func (parcel *Parcel) Delivered() *Product {
	parcel.DeliveredTime = time.Now()
	return parcel.Pdt
}

func main() {

}
