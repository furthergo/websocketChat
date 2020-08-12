package api

import (
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllMessages(c *gin.Context) {
	msgs := models.AllMessages()
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"code": e.SUCCESS,
		"data": msgs,
	})
}

func GetMessageByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	msg, err := models.GetMessageById(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "fail",
			"code": e.ERROR_MESSAGE_NOT_FOUND,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code": e.SUCCESS,
		"data": msg,
	})
}
