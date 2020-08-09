package api

import (
	"fmt"
	"github.com/futhergo/websocketChat/internal/e"
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	r := c.Request
	r.ParseForm()
	fmt.Print(r.Form.Get("username"), r.Form.Get("password"))
	u := models.UserEntity{
		Name: r.Form.Get("username"),
		Password: r.Form.Get("password"),
		Ip: r.RemoteAddr,
	}
	res, _ := u.Auth()
	if !res {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_AUTH,
			"msg": "user auth failed",
		})
		return
	}
	ssid, err := utils.GenerateToken(u.Name, u.Password)
	if err == nil {
		c.SetCookie("session_id", ssid, 1000, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": "login success",
		"data": u.GetMap(),
	})
}

func Register(c *gin.Context) {
	r := c.Request
	r.ParseForm()
	fmt.Print(r.Form.Get("username"), r.Form.Get("password"))
	u := models.UserEntity{
		Name: r.Form.Get("username"),
		Password: r.Form.Get("password"),
		Ip: r.RemoteAddr,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg": "sorry，your input is error, please try again",
		})
		return
	}
	u.Password = string(hash)
	u.Ip = c.Request.RemoteAddr
	u.CreateTime = time.Now()
	u.ModifyTime = time.Now()
	u.Save()
	ssid, err := utils.GenerateToken(u.Name, u.Password)
	if err == nil {
		c.SetCookie("session_id", ssid, 1000, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": "register success",
		"data": u.GetMap(),
	})
}

func LoginFromHtml(c *gin.Context) {
	w := c.Writer
	var u models.UserEntity
	err := c.ShouldBind(&u)
	if err != nil {
		w.Write([]byte("sorry，user is not register"))
	}
	res, _ := u.Auth()
	if !res {
		w.Write([]byte("sorry，user is not register"))
		return
	}
	if err == nil {
		c.SetCookie("session_id", u.Sha256(), 1000, "/", "localhost", false, true)
	}
	c.Redirect(http.StatusPermanentRedirect,  fmt.Sprintf("/chat?username=%v", u.Name))
}

func RegisterFromHtml(c *gin.Context) {
	w := c.Writer
	var u models.UserEntity
	err := c.ShouldBind(&u)
	if err != nil {
		w.Write([]byte("sorry，your input is error"))
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Write([]byte("sorry，your input is error, please try again"))
	}
	u.Password = string(hash)
	u.Ip = c.Request.RemoteAddr
	u.CreateTime = time.Now()
	u.ModifyTime = time.Now()
	u.Save()
	if err == nil {
		c.SetCookie("session_id", u.Sha256(), 1000, "/", "localhost", false, true)
	}
	c.Redirect(http.StatusPermanentRedirect,  fmt.Sprintf("/chat?username=%v", u.Name))
}