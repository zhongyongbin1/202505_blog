package logger

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Logger 包装zerolog.Logger，提供额外功能
type Logger struct {
	*zerolog.Logger
}

// 全局日志记录器
var (
	GlobalLogger Logger
	once         sync.Once
)

type CtxKey string

const (
	loggerCtxKey CtxKey = "logger"
)

// Init 初始化全局日志记录器
func InitLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	level, err := zerolog.ParseLevel("info")
	if err != nil {
		level = zerolog.InfoLevel
	}

	// 创建日志记录器
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(output).Level(level).With().
		Timestamp().
		Str("service", "test-gin").
		Logger()

	// 设置全局日志记录器
	GlobalLogger = Logger{&logger}
}

// Get 获取全局日志记录器
func Get() *Logger {
	once.Do(func() {
		// 只初始化一次
		InitLogger()
	})
	return &GlobalLogger
}

// Info 返回信息级别的日志事件
func Info() *zerolog.Event {
	return Get().Info()
}

// Warn 返回警告级别的日志事件
func Warn() *zerolog.Event {
	return Get().Warn()
}

// Error 返回错误级别的日志事件
func Error() *zerolog.Event {
	return Get().Error()
}

// Debug 返回调试级别的日志事件
func Debug() *zerolog.Event {
	return Get().Debug()
}

// Trace 返回跟踪级别的日志事件
func Trace() *zerolog.Event {
	return Get().Trace()
}

// Fatal 返回致命级别的日志事件
func Fatal() *zerolog.Event {
	return Get().Fatal()
}

// Panic 返回恐慌级别的日志事件
func Panic() *zerolog.Event {
	return Get().Panic()
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	ctx := l.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}

	// 先创建一个命名变量，再获取其地址
	zerologLogger := ctx.Logger()
	newLogger := Logger{&zerologLogger}
	return &newLogger
}

func (l *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerCtxKey, l)
}

// FromContext 从上下文中获取logger
func FromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(loggerCtxKey).(*Logger); ok {
		return l
	}
	return Get() // 默认返回全局logger
}
