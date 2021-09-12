package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 오류가 발생함..

	// 우체국 전송 객체를 만듭니다.
	//sender := &koreaPost.PostSender{}
	//SendBook("어린 왕자", sender)
	//SendBook("그리스인 조르바", sender)
}
