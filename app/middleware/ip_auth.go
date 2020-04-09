package middleware

import (
	"fmt"
	"gin-app-start/app/util"

	"github.com/gin-gonic/gin"
)

func IPAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := util.Context{Ctx: c}
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
			ctx.Response(401, fmt.Sprintf("%s not in ipList", clientIp), nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
