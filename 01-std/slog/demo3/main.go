package main

import (
	"log/slog"

	"github.com/ThreeDotsLabs/watermill"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
1.记录到文件
2.使用 lumberjack 进行日志切割
*/

func main() {
	r := &lumberjack.Logger{
		Filename:   "./01-std/slog/demo3/foo.log",
		LocalTime:  true,
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 5,
		Compress:   false,
	}

	var lvl slog.LevelVar
	lvl.Set(slog.LevelInfo)
	l := slog.New(slog.NewTextHandler(r, &slog.HandlerOptions{
		Level: &lvl,
	}))
	watermill.NewSlogLogger(l)
	slog.SetDefault(l)

	for i := 0; i < 100000; i++ {
		slog.Info("greeting", "say", "hello")
	}
}
