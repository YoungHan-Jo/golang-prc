package main

import (
	"goproject/ch20/fedex"
	koreapost "goproject/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}

	SendBook("clean architecture", sender)
	SendBook("clean code", sender)

	sender2 := &koreapost.PostSender{}
	SendBook("clean architecture", sender2)
	SendBook("clean code", sender2)
}
