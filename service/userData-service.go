package service

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
)

func userData() {

	wsHandler := func(message []byte) {
		fmt.Println(string(message))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsUserDataServe(listenKey, wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
