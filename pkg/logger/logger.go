package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func Init(env string, logFile string) error {
	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	if logFile != "" {
		// Ensure the log directory exists
		logDir := filepath.Dir(logFile)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return err
		}

		config.OutputPaths = []string{"stdout", logFile}
		config.ErrorOutputPaths = []string{"stderr", logFile}
	}

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	globalLogger = logger
	return nil
}

func GetLogger() *zap.Logger {
	if globalLogger == nil {
		logger, _ := zap.NewDevelopment()
		globalLogger = logger
	}
	return globalLogger
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

func With(fields ...zap.Field) *zap.Logger {
	return GetLogger().With(fields...)
}

func WithContext(fields ...zap.Field) *zap.Logger {
	return GetLogger().With(fields...)
}

func Close() {
	if globalLogger != nil {
		_ = globalLogger.Sync()
	}
}

func InitDefault() {
	if globalLogger == nil {
		logger, _ := zap.NewDevelopment()
		globalLogger = logger
	}
}

func init() {
	if globalLogger == nil {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err := config.Build(zap.AddCallerSkip(1))
		if err != nil {
			logger, _ = zap.NewDevelopment()
		}
		globalLogger = logger
	}
}

