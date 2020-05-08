package controller

import (
	"gin-app-start/app/common"
	"gin-app-start/app/schema"
	"gin-app-start/app/util"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	ctx := util.Context{Ctx: c}

	var user schema.User
	// user := &schema.User{}
	if err := ctx.Validate(&user); err != nil {
		return
	}

	// form-data
	// name := c.PostForm("name")
	// password := c.PostForm("password")

	// json
	name := user.Name
	password := user.Password
	log.Println("name:", name, "password:", password)

	// todo从db验证
	if name == "admin" && password == "admin" {
		session.Set("user_name", name)
		session.Save()

		content := map[string]string{
			"data": "success",
		}
		ctx.Response(0, nil, content)
	} else {
		ctx.Response(401, common.LOGIN_FAIL, nil)
	}
	return
}
