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

func LoginFromHtml(c *gin.Context) {
	w := c.Writer
	var u models.UserEntity
	err := c.ShouldBind(&u)
	if err != nil {
		w.Write([]byte("sorry，user is not allowed"))
	}
	res, _ := u.Auth()
	if !res {
		w.Write([]byte("sorry，user is not allowed"))
		return
	}
	c.Redirect(http.StatusPermanentRedirect,  fmt.Sprintf("/chat?username=%v&password=%v", u.Name, u.Password))
}

func RegisterFromHtml(c *gin.Context) {

}