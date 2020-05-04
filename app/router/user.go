package router

import (
	"gin-app-start/app/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	router := group.Group("")
	{
		// 健康检查
		router.POST("/login", controller.Login)
		// router.POST("/lgout", controller.Lgout)
	}
}
