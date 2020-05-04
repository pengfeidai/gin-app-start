package router

import (
	"fmt"
	"gin-app-start/app/config"

	"gin-app-start/app/middleware"
	"gin-app-start/app/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(config.Mode)
	router := gin.Default()
	config := config.Conf
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
	// router.Use(middleware.IPAuth())
	// 限流
	router.Use(middleware.Limit(config.LimitNum))

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))

	// 路由分组加载
	group := router.Group(config.Url.Prefix)
	InitHealthCheckRouter(group)
	InitUserRouter(group)

	// user
	return router
}
