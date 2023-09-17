package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func diningProblem(name string, first, second *sync.Mutex, firstName, secendName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s trying to eat rice.\n", name)
		first.Lock()
		fmt.Printf("%s %s get.\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s get.\n", name, secendName)

		fmt.Printf("%s eats rice.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "fork", "spoon")
	go diningProblem("B", spoon, fork, "spoon", "fork")
	wg.Wait()

}
