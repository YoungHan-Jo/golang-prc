package main

import (
	"ch16/2/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI: ", publicpkg.PI)
	publicpkg.PublicFunc()

	var myInt publicpkg.MyInt = 1
	fmt.Println(myInt)

	var mystruct = publicpkg.MyStruct{Age: 10}
	fmt.Println("mystruct:", mystruct)

}
