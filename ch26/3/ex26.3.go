package main

import "fmt"

type T interface {
	Something()
}

type S struct {
}

func (s *S) Something() {
	fmt.Println("somthing s")
}

type U struct {
}

func (u *U) Something() {
	fmt.Println("somthing u")
}

func q(t T) {
	t.Something()
}

type Z struct {
}

func (z *Z) Anything() {
	fmt.Println("anything z")
}

var y = &S{}
var u = &U{}
var z = &Z{}

func main() {
	q(y)
	q(u)
	// q(z)
}
