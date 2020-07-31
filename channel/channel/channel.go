/**
channel是一等公民:
	作为数组
	作为参数
	作为返回值
	单向channel
	buffer channel
	close channel, close永远只能是发送端, 通知接收端数据发完了
*/

package main

import (
	"fmt"
	"time"
)

func worker(i int, ch chan int) {
	for n := range ch {
		fmt.Printf("worker %d recive %d\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	ch := make(chan int)
	go worker(i, ch)
	return ch
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func channelClose() {
	ch := make(chan int)
	go worker(0, ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	time.Sleep(time.Millisecond)
}

func main() {
	//channelDemo()
	channelClose()
}
