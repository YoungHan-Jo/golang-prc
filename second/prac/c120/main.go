package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

type Box struct {
	flows *ring.Ring
}

func (b *Box) push(sc *bufio.Scanner) {
	sc.Scan()
	runes := []rune(sc.Text())

	for _, v := range runes {
		b.flows.Value = string(v)
		b.flows = b.flows.Next()
	}
}

func (b *Box) turn() *Box {
	b.flows = b.flows.Next()
	return b
}

func (b *Box) view() string {
	var value = ""
	for i := 0; i < b.flows.Len(); i++ {
		value += b.flows.Value.(string)
		b.flows = b.flows.Next()
	}
	return value
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	boxA := Box{ring.New(n)}
	boxB := Box{ring.New(n)}

	boxA.push(sc)
	boxB.push(sc)

	result := "No"

	for i := 0; i < n+1; i++ {
		if boxA.view() == boxB.view() {
			result = "Yes"
			break
		}
		boxA.turn()
	}

	fmt.Println(result)
}
