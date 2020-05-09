package router

import (
	"fmt"
	"gin-app-start/app/config"

	"gin-app-start/app/middleware"
	"gin-app-start/app/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	config := config.Conf
	gin.SetMode(gin.TestMode)
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
	router.Use(middleware.Logger())
	// ip白名单
	// router.Use(middleware.IPAuth())
	// 限流
	router.Use(middleware.Limit(config.LimitNum))

	var store sessions.Store
	if config.Server.UserRedis {
		store, _ = redis.NewStore(config.Session.Size, "tcp", config.Redis.Addr, config.Redis.Password, []byte("secret"))
	} else {
		store = cookie.NewStore([]byte("secret"))
	}
	// store.Options(sessions.Options{
	// 	HttpOnly: true,
	// 	MaxAge:   60 * 15,
	// })
	router.Use(sessions.Sessions("session_id", store))

	// 路由分组加载
	group := router.Group(config.Url.Prefix)
	InitHealthCheckRouter(group)
	InitUserRouter(group)

	// user
	return router
}
