package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Trade struct {
	Exchange  string  `json:"exchange"`
	Base      string  `json:"base"`
	Quote     string  `json:"quote"`
	Direction string  `json:"direction"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
	PriceUsd  float64 `json:"priceUsd"`
}

func StreamCoinCap(context *gin.Context) {
	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var jsonT []byte
	// websocket client connection
	c, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/trades/binance", nil)
	if err != nil {
		panic(err)
	}

	// producer: read data stream from websocket and send to channel
	input := make(chan Trade) // 1️⃣

	go func() { // 2️⃣
		// read from the websocket
		for {
			_, message, err := c.ReadMessage() // 3️⃣
			if err != nil {
				break
			}
			// unmarshal the message
			var trade Trade
			json.Unmarshal(message, &trade) // 4️⃣
			// send the trade to the channel
			input <- trade // 5️⃣
		}
		close(input) // 6️⃣
	}()

	dogecoin := make(chan Trade)
	go func() {
		for trade := range input {
			if trade.Base == "bitcoin" && trade.Quote == "tether" {
				dogecoin <- trade
			}
		}
		close(dogecoin)
	}()
	//jsonArr := make(chan string)
	// print the trades
	for trade := range dogecoin {
		jsonT, _ = json.Marshal(trade)
		//jsonArr <- string(jsonT)

		err = ws.WriteMessage(1, jsonT)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	defer c.Close()
}
