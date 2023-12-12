package main

import (
	"context"
	"log/slog"
	"net"
	"os"
)

func main() {
	var lvl slog.LevelVar
	lvl.Set(slog.LevelDebug)
	l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: &lvl,
	}))
	slog.SetDefault(l)

	slog.Info("before resetting log level:")
	slog.Debug("greeting", "name", "tony")
	slog.Info("greeting", "name", "tony")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	slog.Info("after resetting log level to error level:")
	lvl.Set(slog.LevelError)
	slog.Info("greeting", "name", "tony")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
