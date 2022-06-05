package service

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func UserWallet(*gin.Context) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("API_SECRET")
	client := binance.NewClient(apiKey, secretKey)

	var result []binance.Balance
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	result = res.Balances
	fmt.Println(result[0])
}
