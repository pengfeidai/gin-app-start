package middleware

import (
	"time"

	"gin-app-start/app/common"
	"gin-app-start/app/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if uri == "/favicon.ico" {
			return
		}
		//开始时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		method := c.Request.Method
		statusCode := c.Writer.Status()
		ip := c.ClientIP()
		if !config.Conf.Log.Debug {
			common.Logger.WithFields(logrus.Fields{
				"clientIp":    ip,
				"statusCode":  statusCode,
				"reqMethod":   method,
				"reqUri":      uri,
				"latencyTime": latencyTime,
			}).Info()
		} else {
			now := time.Now().Format(common.TIME_FORMAT)
			common.Logger.Infof("%s | %3d | %13v | %15s | %s  %s",
				now,
				statusCode,
				latencyTime,
				ip,
				method,
				uri,
			)
		}
	}
}
