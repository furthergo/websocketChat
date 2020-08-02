package api

import (
	"fmt"
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	r := c.Request
	w := c.Writer
	r.ParseForm()
	fmt.Print(r.Form.Get("username"), r.Form.Get("password"))
	u := models.UserEntity{
		Name: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}
	res, _ := u.Auth()
	if !res {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("sorry，user is not allowed"))
		return
	}
}