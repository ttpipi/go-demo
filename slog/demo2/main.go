package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"
)

/*
自定义后端实现
*/

func main() {
	var ch = make(chan []byte, 100)
	arrts := []slog.Attr{
		{
			Key:   "field1",
			Value: slog.StringValue("value1"),
		},
		{
			Key:   "field2",
			Value: slog.StringValue("value2"),
		},
	}
	slog.SetDefault(slog.New(NewChanHandler(ch).WithAttrs(arrts)))

	go func() {
		for {
			b := <-ch
			fmt.Println(string(b))
		}
	}()

	slog.Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	time.Sleep(3 * time.Second)
}
