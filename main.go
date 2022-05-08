package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"golang.org/x/net/context"
)

var (
	apiKey    = "wdA1ZXkD2VgF7p3v1vgI8CkgsXXoT1IqFCTeo8XAWDbAnitoFry4Z53iCqCJsBhU"
	secretKey = "6RysiPFMfRTYOA3mN5JcfUwQLKfM7YD4QglmIsoH1W5OwQWQgEtRd54JssQQ1fAA"
)
var client = binance.NewClient(apiKey, secretKey)
var futuresClient = binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
var deliveryClient = binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

func startUserStream() {
	res, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func userWallet() {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

//func userData() {
//	wsHandler := func(message []byte) {
//		fmt.Println(string(message))
//	}
//	errHandler := func(err error) {
//		fmt.Println(err)
//	}
//	doneC, _, err := binance.WsUserDataServe(listenKey, wsHandler, errHandler)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	<-doneC
//}
func main() {

	userWallet()
}
