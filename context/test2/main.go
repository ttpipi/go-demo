/**
使用Contex控制多个子协程, 多个子协程都会收到退出
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[监控1]")
	go watch(ctx, "[监控2]")
	go watch(ctx, "[监控3]")

	time.Sleep(10 * time.Second)
	fmt.Println("通用在所有子协程退出")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "子协程退出")
			return
		default:
			fmt.Println(name, "子协程监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
