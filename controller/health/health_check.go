package health

import (
	"fmt"
	"gin-app-start/middleware"
	"gin-app-start/schema"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}
	// 参数校验
	// var person Person
	h := &schema.Health{}
	if err := ctx.Validate(h); err != nil {
		return
	}

	name := c.Query("name")
	content := map[string]string{
		"data": fmt.Sprintf("hello %s", name),
	}
	ctx.Response(nil, content)
	return
}
