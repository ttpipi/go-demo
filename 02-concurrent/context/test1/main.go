/**
使用Contex控制一个子协程
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子协程退出")
				return
			default:
				fmt.Println("子协程工作中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("通知子协程退出")
	cancel()
	time.Sleep(5 * time.Second)
}
