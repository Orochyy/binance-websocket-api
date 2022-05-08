package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type wallet struct {
	BTC float64 `json:"btc"`
}

var (
	apiKey    = "wdA1ZXkD2VgF7p3v1vgI8CkgsXXoT1IqFCTeo8XAWDbAnitoFry4Z53iCqCJsBhU"
	secretKey = "6RysiPFMfRTYOA3mN5JcfUwQLKfM7YD4QglmIsoH1W5OwQWQgEtRd54JssQQ1fAA"
)
var client = binance.NewClient(apiKey, secretKey)

//func startUserStream() {
//	res, err := client.NewStartUserStreamService().Do(context.Background())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(res)
//}

func userWalletBTC(*gin.Context) {
	var result []binance.Balance
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	result = res.Balances
	fmt.Println(result[0])
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/wallet", userWalletBTC)
	r.Run(":8080")

}
