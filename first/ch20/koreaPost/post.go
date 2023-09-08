package koreapost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("post sends %v parcel\n", parcel)
}
