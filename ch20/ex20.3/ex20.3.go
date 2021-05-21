package main

import (
	"goproject/ch20/ex20.3/fedex"
	"goproject/ch20/ex20.3/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린왕자", sender)
	SendBook("그리스인 조르바", sender)
}