package router

import (
	"fmt"
	"gin-app-start/app/config"

	"gin-app-start/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	config := config.Conf
	gin.SetMode(config.Mode)
	router := gin.New()
	// 404处理
	router.NoRoute(func(c *gin.Context) {
		ctx := middleware.Context{Ctx: c}
		path := c.Request.URL.Path
		method := c.Request.Method
		ctx.Response(404, fmt.Sprintf("%s %s not found", method, path), nil)
	})

	// 中间件
	router.Use(
		cors.Default(),
		middleware.Recovery(),
		middleware.Logger(),
	)

	var store sessions.Store
	if config.Server.UserRedis {
		store, _ = redis.NewStore(config.Session.Size, "tcp", config.Redis.Addr, config.Redis.Password, []byte("secret"))
	} else {
		store = cookie.NewStore([]byte("secret"))
	}

	store.Options(sessions.Options{
		Path:     config.Session.Path,
		HttpOnly: config.Session.HttpOnly,
		MaxAge:   config.Session.MaxAge,
	})

	router.Use(sessions.Sessions("session_id", store))

	// 路由分组加载
	group := router.Group(config.Url.Prefix)
	InitHealthCheckRouter(group)

	// user
	return router
}
