package router

import (
	"gin-app-start/app/controller"

	"github.com/gin-gonic/gin"
)

func InitHealthCheckRouter(group *gin.RouterGroup) {
	router := group.Group("")
	// .Use(middleware.IPAuth())
	{
		// 健康检查
		router.GET("/health_check", controller.CheckHealth)
	}
}
