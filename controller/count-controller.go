package controller

import (
	"binance-websocket-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CountController interface {
	GetCount(c *gin.Context)
}

func GetCount(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.Sum)
}
