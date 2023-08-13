package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	// a := Actor{name, hp, speed}
	// return &a

	a := new(Actor)
	a.Name = name
	a.HP = hp
	a.Speed = speed

	return a
}

func main() {
	fmt.Println(NewActor("AAA", 10, 0.4))
}
