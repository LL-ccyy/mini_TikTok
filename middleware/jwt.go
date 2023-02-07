package middleware

import (
	"Minimalist_TikTok/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 0
		//token := c.GetHeader("Authorization")
		token := c.Query("token")
		fmt.Println("token=", token)
		if token == "" { //无token
			code = 1
		} else {
			claims, err := util.ParseToken(token)
			if err != nil { //token无权限，错误
				code = 1
			} else if time.Now().Unix() > claims.ExpiresAt { //token失效
				code = 1
			}
		}
		if code == 1 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "User doesn't exist",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
