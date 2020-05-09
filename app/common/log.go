package common

import (
	"bufio"
	"gin-app-start/app/config"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	config := config.Conf.Log

	// 生产环境写入文件JSON格式
	if !config.Debug {
		fileName := path.Join(config.DirName, config.FileName)
		Logger.SetFormatter(&logrus.JSONFormatter{})

		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		// file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			panic(err)
		}
		//设置日志级别
		Logger.SetLevel(logrus.InfoLevel)
		//设置输出
		Logger.SetOutput(bufio.NewWriter(file))

		// 设置 rotatelogs
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名称
			fileName+".%Y-%m-%d.log",
			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(fileName),
			// 设置最大保存时间(15天)
			rotatelogs.WithMaxAge(config.MaxAge*24*time.Hour),
			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}))
	} else {
		Logger.SetLevel(logrus.DebugLevel)
	}
}
