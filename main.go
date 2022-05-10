package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"os"
	"time"
)

var (
	apiKey    = os.Getenv("API_KEY")
	secretKey = os.Getenv("API_SECRET")
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

func userWallet(*gin.Context) {
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

	countAVGBTC(37400, 35711, 36083, 31973, 4150, 3076, 1490, 3236)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.POST("/count", countAVGBTC(37400, 35711, 36083, 31973, 4150, 3076, 1490, 3236))
	r.GET("/wallet", userWallet)
	r.Run(":8080")
}

func countAVGBTC(btcPrice1 float64, btcPrice2 float64, btcPrice3 float64, btcPrice4 float64, btcUSD1 float64, btcUSD2 float64, btcUSD3 float64, btcUSD4 float64) {
	allMoney := btcUSD1 + btcUSD2 + btcUSD3 + btcUSD4
	USDresult := fmt.Sprintf("all deposit:%v USDT", allMoney)
	fmt.Println(USDresult)
	time.Sleep(time.Millisecond * 500)
	btcUSD1Percent1 := (btcUSD1 * 100) / allMoney
	btcUSD1Percent2 := (btcUSD2 * 100) / allMoney
	btcUSD1Percent3 := (btcUSD3 * 100) / allMoney
	btcUSD1Percent4 := (btcUSD4 * 100) / allMoney
	allPercents := btcUSD1Percent1 + btcUSD1Percent2 + btcUSD1Percent3 + btcUSD1Percent4
	if allPercents != 100 {
		panic("Panic: allPercents != 100% ")
		time.Sleep(time.Second)
	}
	result := ((btcPrice1 * btcUSD1) + (btcPrice2 * btcUSD2) + (btcPrice3 * btcUSD3) + (btcPrice4 * btcUSD4)) / allMoney
	fmt.Println(result)
}
