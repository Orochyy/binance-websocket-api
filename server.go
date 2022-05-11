package main

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"os"
	"time"
)

const sellPrice = 40000

var BTCPrice = map[int]float64{
	1: 31000,
	2: 30450,
	//3: 36083,
	//4: 31973,
	//5: 31000,
}
var btcUSD = map[int]float64{
	1: 10670,
	2: 993,
	//3: 1490,
	//4: 3236,
	//5: 850,
}

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
func userWallet(*gin.Context) {
	var result []binance.Balance
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	result = res.Balances
	fmt.Println(result[0])
}

func main() {
	//AVGBuy := countAVGBTC(BTCPrice, btcUSD)
	//PercentDiff := countDiff(sellPrice, AVGBuy)
	//time.Sleep(time.Millisecond * 300)
	//CountPl := countPlus(PercentDiff, allMoney(btcUSD))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/count", sum)
	r.GET("/wallet", userWallet)
	r.Run(":8080")
}

func countAVGBTC(BTCPrice map[int]float64, btcUSD map[int]float64) float64 {
	var allPercents float64
	var AVG float64

	USDresult := fmt.Sprintf("all deposit:%v USDT", allMoney(btcUSD))
	fmt.Println(USDresult)
	time.Sleep(time.Millisecond * 500)
	for _, value := range btcUSD {
		var btcUSDPercent = (value * 100) / allMoney(btcUSD)
		allPercents += btcUSDPercent
	}
	if allPercents != 100 {
		panic("Panic: allPercents != 100% ")
		time.Sleep(time.Second)
	}
	for key := range BTCPrice {
		AVG += BTCPrice[key] * btcUSD[key]
	}
	AVG /= allMoney(btcUSD)
	resultSf := fmt.Sprintf("AVG USDT BUY PRICE:%v USD", int(AVG))
	fmt.Println(resultSf)
	time.Sleep(time.Millisecond * 500)
	return AVG
}

func countDiff(sellPrice float64, buyPrice float64) float64 {
	percents := (sellPrice * 100) / buyPrice
	result := percents - 100
	resultSf := fmt.Sprintf("%.2f %%", percents-100)
	fmt.Println(resultSf)
	return result
}

func countPlus(percent float64, deposit float64) float64 {
	result := (deposit * percent) / 100
	resultSf := fmt.Sprintf("+%.2f USD", result)
	fmt.Println(resultSf)
	return result
}
func allMoney(btcUSD map[int]float64) float64 {
	var allMoney float64
	for _, x := range btcUSD {
		allMoney += x
	}
	return allMoney
}
func sum(*gin.Context) {
	AVGBuy := countAVGBTC(BTCPrice, btcUSD)
	PercentDiff := countDiff(sellPrice, AVGBuy)
	time.Sleep(time.Millisecond * 300)
	CountPl := countPlus(PercentDiff, allMoney(btcUSD))

	result := allMoney(btcUSD) + CountPl
	time.Sleep(time.Millisecond * 300)
	resultSf := fmt.Sprintf("%.2f USD", result)
	fmt.Println(resultSf)
}
