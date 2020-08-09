package main

import (
	"github.com/futhergo/websocketChat/internal/DB"
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/router/api/v1"
	logger "github.com/futhergo/websocketChat/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	logger.Log2file("~~~~~~~~~~New Server Start~~~~~~~~~")

	DB.InitDB()

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
	idx.Any("/chat", /*middleware.Cookies(), */func(c *gin.Context) {
		u := models.UserEntity{
			Name: c.Query("username"),
		}
		c.SetCookie("session_id", u.Sha256(), 1000, "/", ":", false, true)
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/user/login", api.Login)
		apiV1.GET("/user/resister", api.Register)
		apiV1.POST("/loginFromHTML", api.LoginFromHtml)
		apiV1.POST("/registerFromHTML", api.RegisterFromHtml)
	}

	wsS := models.NewWsChatServer()
	router.GET("/ws/msg", wsS.MessageAPI)

	//start http server
	router.Run(":1111")
}