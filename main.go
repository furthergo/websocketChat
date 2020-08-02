package main

import (
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/router/api/v1"
	logger "github.com/futhergo/websocketChat/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	logger.Log2file("~~~~~~~~~~New Server Start~~~~~~~~~")

	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.Static("/statics", "./statics")
	idx := router.Group("/")
	idx.Any("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to my WebSocket Chat~~",
			"subtitle": "You can use name `test` and password `test` to login, have a good day!",
		})
	})
	idx.POST("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})
	// routers
	router.GET("/api/v1/login", api.Login)
	router.POST("/api/v1/loginFromHTML", api.LoginWithRedirect)

	wsS := models.NewWsChatServer()
	router.GET("/ws/msg", wsS.MessageAPI)

	//start http server
	router.Run(":80")
}