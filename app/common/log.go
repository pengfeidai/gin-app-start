package common

import (
	"gin-app-start/app/config"
	"os"

	"github.com/sirupsen/logrus"
)

var ErrorLog *logrus.Logger
var AccessLog *logrus.Logger

func InitLog() {
	initErrorLog()
	initAccessLog()
}

func initErrorLog() {
	ErrorLog = logrus.New()
	errorLogFile := config.Conf.Log.ErrorLogFile
	// 设置日志格式为json格式
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile(errorLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	ErrorLog.SetOutput(file)
}

func initAccessLog() {
	AccessLog = logrus.New()
	AccessLogFile := config.Conf.Log.AccessLogFile
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile(AccessLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	AccessLog.SetOutput(file)
}
