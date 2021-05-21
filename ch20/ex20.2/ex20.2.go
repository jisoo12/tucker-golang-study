package main

import "goproject/ch20/ex20.2/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("그리스인 조르바", sender)
}