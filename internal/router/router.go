package router

import (
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/router/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.Static("/statics", "./statics")
	idx := router.Group("/")
	idx.Any("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to my WebSocket Chat~~",
			"subtitle": "Have a good day!（haocheng SB）",
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
		apiUser := apiV1.Group("/user")
		{
			apiUser.GET("/login", api.Login)
			apiUser.GET("/resister", api.Register)
			apiUser.POST("/loginFromHTML", api.LoginFromHtml)
			apiUser.POST("/registerFromHTML", api.RegisterFromHtml)
		}
		apiV1.GET("/allMessages", api.GetAllMessages)
		apiV1.GET("/message", api.GetMessageByID)
	}

	wsS := models.NewWsChatServer()
	router.GET("/ws/msg", wsS.MessageAPI)

	//start http server
	router.Run(":1111")
}