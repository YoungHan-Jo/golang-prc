package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)

	fmt.Println(data.value)
	fmt.Println(data.data[100])

	var p *Data = &Data{}

	fmt.Println(p)

	var p2 = new(Data)
	p2.data[0] = 100
	p2.value = 1
	fmt.Println("p2:", p2)

}
