package main

import "fmt"

const Pig int = 0
const Cat int = 1
const Dog int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("pipipipipipig")
	} else if animal == Cat {
		fmt.Println("cacacacacacat")
	} else if animal == Dog {
		fmt.Println("dododododododog")
	} else {
		fmt.Println("i don't know")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cat)
	PrintAnimal(Dog)
}
