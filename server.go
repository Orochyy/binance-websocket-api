package main

import (
	"binance-websocket-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func startUserStream() {
//	res, err := client.NewStartUserStreamService().Do(context.Background())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(res)
//}

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
//func userWallet(*gin.Context) {
//	var result []binance.Balance
//	res, err := client.NewGetAccountService().Do(context.Background())
//	if err != nil {
//		fmt.Println(err)
//	}
//	result = res.Balances
//	fmt.Println(result[0])
//
//}

func main() {

	ip := string(service.LocalIP())
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Login",
			},
		)
	})
	r.GET("/count", service.Sum)
	r.GET("/wallet", service.UserWallet)
	r.GET("/stream", service.StreamCoinCap) //handle the error so that it doesn't drop when I log in from the browser
	r.Run(ip + ":8080")
}
