package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parseRandomJSON-test/service"
)

type CountController interface {
	GetCount(c *gin.Context)
}

func GetCount(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ParseJokeNameStruct())
}
