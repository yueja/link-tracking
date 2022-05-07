package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	// Type 日志类型
	Type = "app"
)

type LoggerConfig struct {
	AppName string `json:"app_name"`
}

// Log 应用程序日志句柄
var (
	log     *logrus.Logger
	appName string
)

// NewLogger 创建新的logger
func (config LoggerConfig) NewLogger() *logrus.Logger {
	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = os.Stdout

	// 设置日志级别
	logger.SetLevel(logrus.TraceLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	appName = config.AppName
	return logger
}

func Init(config LoggerConfig) {
	log = config.NewLogger()
}

// PanicW panic日志
func PanicW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.PanicLevel, template, args)
}

// FatalW fatal日志
func FatalW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.FatalLevel, template, args)
}

// ErrorW error错误日志
func ErrorW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.ErrorLevel, template, args)
}

// WarnW warn日志
func WarnW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.WarnLevel, template, args)
}

// InfoW info日志
func InfoW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.InfoLevel, template, args)
}

// DebugW debug日志
func DebugW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.DebugLevel, template, args)
}

// TraceW trace日志
func TraceW(ctx context.Context, template string, args ...interface{}) {
	newServer().setLog(ctx, logrus.TraceLevel, template, args)
}
