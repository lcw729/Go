package main

import (
	// "Go/practice/ch20/fedex"
	"Go/practice/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

// 번거롭게 기존 코드를 바꿔야한다. [산탄총 코드] => interface가 필요하다.
// 꼭 필요한 함수를 가지고 있기만 하면 모두 적용가능
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &koreaPost.PostSender{}

	SendBook("어린 왕자", sender)
	SendBook("돼지 삼형제", sender)
}
