package main

import (
	"binance-websocket-api/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {float64} countSum
// @Router /count [get]

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(ip + ":8080")
}
