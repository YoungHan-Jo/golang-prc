package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < 5; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}

	sc = bufio.NewScanner(os.Stdin)

	sc.Split(bufio.ScanLines)
	sc.Scan()
	fmt.Println(sc.Text())
	sc.Scan()
	fmt.Println(sc.Text())

}
