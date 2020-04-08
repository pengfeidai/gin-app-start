package user

import (
	"gin-app-start/middleware"
	"gin-app-start/schema"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	ctx := middleware.Context{Ctx: c}

	// var user schema.User
	user := &schema.User{}
	if err := ctx.Validate(user); err != nil {
		return
	}

	// form-data
	// name := c.PostForm("name")
	// password := c.PostForm("password")
	name := user.Name
	password := user.Password
	log.Println("name:", name, "password:", password)

	// todo从db验证
	if name == "admin" && password == "admin" {
		session := sessions.Default(c)
		session.Set("user_name", name)
		session.Save()
		content := map[string]string{
			"data": "success",
		}
		ctx.Response(nil, content)
	} else {
		ctx.Response("账号或密码错误", nil)
	}
	return
}
