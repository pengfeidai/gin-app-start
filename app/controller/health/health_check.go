package health

import (
	"fmt"
	"gin-app-start/app/schema"
	"gin-app-start/app/util"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	ctx := util.Context{Ctx: c}
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
	ctx.Response(0, nil, content)
	return
}
