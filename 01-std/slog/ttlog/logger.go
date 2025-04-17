package ttlog

import (
	"context"
	"sync/atomic"
)

var defaultLogger atomic.Pointer[Logger]

func init() {
	defaultLogger.Store(New(newDefaultHandler()))
}

func Default() *Logger { return defaultLogger.Load() }

type Logger struct {
	handler Handler
}

func New(h Handler) *Logger {
	if h == nil {
		panic("nil Hanlder")
	}
	return &Logger{handler: h}
}

// ttlin: 为什么要提供public方法
// func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any) {
// 	l.log(ctx, level, msg, args...)
// }

func (l *Logger) log(ctx context.Context, level Level, msg string, args ...any) {
	l.handler.Handle(ctx, level)
}

// ttlin:
// 提供包级别的函数
// 必须有一个Logger对象
func Log(ctx context.Context, level Level, msg string, args ...any) {
	Default().log(ctx, level, msg, args...)
}
