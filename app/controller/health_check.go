package controller

import (
	"fmt"
	"gin-app-start/app/middleware"
	"gin-app-start/app/schema"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	// 参数校验
	var p schema.Health
	if err := ctx.Validate(&p); err != nil {
		return
	}

	name := c.Query("name")
	content := map[string]string{
		"data": fmt.Sprintf("hello %s", name),
	}
	ctx.Response(0, nil, content)
	return
}
