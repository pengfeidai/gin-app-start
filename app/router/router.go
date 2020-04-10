package router

import (
	"fmt"
	"gin-app-start/app/controller"
	"gin-app-start/app/middleware"
	"gin-app-start/app/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(config.Mode)
	router := gin.Default()

	// 404处理
	router.NoRoute(func(c *gin.Context) {
		ctx := util.Context{Ctx: c}
		path := c.Request.URL.Path
		method := c.Request.Method
		ctx.Response(404, fmt.Sprintf("%s %s not found", method, path), nil)
	})

	// 跨域
	router.Use(cors.Default())
	// ip白名单
	router.Use(middleware.IPAuth())

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))

	// 登录登出
	router.POST("/login", controller.Login)
	// router.POST("/lgout", user.Lgout)

	// 健康检查
	router.GET("/health_check", controller.CheckHealth)

	// user
	return router
}
