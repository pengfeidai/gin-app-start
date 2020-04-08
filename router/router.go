package router

import (
	"gin-app-start/controller/health"
	"gin-app-start/controller/user"
	"gin-app-start/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 跨域
	router.Use(cors.Default())
	// ip白名单
	router.Use(middleware.IPAuth())

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))

	// 登录登出
	router.POST("/login", user.Login)
	// router.POST("/lgout", user.Lgout)

	// 健康检查
	router.GET("/health_check", health.CheckHealth)

	// user
	return router
}
