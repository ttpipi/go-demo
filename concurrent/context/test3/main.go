/**
WithValue传递元数据
*/

package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctxValue := context.WithValue(ctx, key, "[监控1]")
	go watch(ctxValue)

	time.Sleep(10 * time.Second)
	fmt.Println("通知子协程退出")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "子协程退出")
			return
		default:
			fmt.Println(ctx.Value(key), "子协程监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
