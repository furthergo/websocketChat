package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cookies() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookies, e := c.Request.Cookie("session_id")
		if e == nil {
			c.SetCookie(cookies.Name, cookies.Value, 1000, cookies.Path, cookies.Domain, false, true)
			c.Next()
		} else {
			c.Abort()
			c.Redirect(http.StatusPermanentRedirect, "")
		}
	}
}