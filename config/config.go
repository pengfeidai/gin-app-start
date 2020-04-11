package config

const (
	Port = ":9060"
	Mode = "release"

	ReadTimeout  = 120000
	WriteTimeout = 120000

	LimitNum = 20

	// mongo
	MongoUrl = "mongodb://127.0.0.1:27017"

	// 日志文件
	AccessLogName = "logs/access.log"
	ErrorLogName  = "logs/error.log"
)
