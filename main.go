package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	hub := newHub()
	go hub.run()
	r.GET("/", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
