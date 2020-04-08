package middleware

import (
	"github.com/gin-gonic/gin"
)

func IPAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context{Ctx: c}
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIp := c.ClientIP()
		for _, value := range ipList {
			if clientIp == value {
				flag = true
				break
			}
		}
		if !flag {
			ctx.Response(clientIp+" not in ipList", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
