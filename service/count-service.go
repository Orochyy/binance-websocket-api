package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const sellPrice = 40000
const separate = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"

var BTCPrice = map[int]float64{
	1: 30623,
	//2: 30450,
	//3: 36083,
	//4: 31973,
	//5: 31000,
}
var btcUSD = map[int]float64{
	1: 11950,
	//2: 993,
	//3: 1490,
	//4: 3236,
	//5: 850,
}

func countAVGBTC(BTCPrice map[int]float64, btcUSD map[int]float64) float64 {
	var allPercents float64
	var AVG float64

	USDresult := fmt.Sprintf("%s all deposit:%v USDT", separate, allMoney(btcUSD))
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
	resultSf := fmt.Sprintf("%sAVG USDT BUY PRICE:%v USD", separate, int(AVG))
	fmt.Println(resultSf)
	time.Sleep(time.Millisecond * 500)
	return AVG
}

func countDiff(sellPrice float64, buyPrice float64) float64 {
	percents := (sellPrice * 100) / buyPrice
	result := percents - 100
	resultSf := fmt.Sprintf("%s%.2f %%", separate, percents-100)
	fmt.Println(resultSf)
	return result
}

func countPlus(percent float64, deposit float64) float64 {
	result := (deposit * percent) / 100
	resultSf := fmt.Sprintf("%s +%.2f USD", separate, result)
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
func Sum(*gin.Context) {
	AVGBuy := countAVGBTC(BTCPrice, btcUSD)
	PercentDiff := countDiff(sellPrice, AVGBuy)
	time.Sleep(time.Millisecond * 300)
	CountPl := countPlus(PercentDiff, allMoney(btcUSD))

	result := allMoney(btcUSD) + CountPl
	time.Sleep(time.Millisecond * 300)
	resultSf := fmt.Sprintf("%s%.2f USD", separate, result)
	fmt.Println(resultSf)
}
