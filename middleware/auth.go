package middleware

import (
	"fmt"
	"net/http"
	"oj/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		fmt.Println(uid)
		if uid != nil {
			c.Set("UserID", uid.(uint))
			fmt.Println("CurrentUser is ", uid.(uint))
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("UserID"); ok {
			c.Next()
		}
		c.JSON(http.StatusForbidden, serializer.CheckLogin())
		c.Abort()
	}
}
