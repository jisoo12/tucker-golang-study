package main

import (
	"goproject/ch20/ex20.4/fedex"
	"goproject/ch20/ex20.4/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}