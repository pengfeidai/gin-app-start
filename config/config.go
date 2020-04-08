package config

const (
	Port = ":9060"
	Mode = "release"

	ReadTimeout  = 120000
	WriteTimeout = 120000

	// 日志文件
	AccessLogName = "log/access.log"
	ErrorLogName  = "log/error.log"
)
