/**
解决主程序发完数据后退出, 所有goroutine都被杀掉的问题
也就是说goroutine收完数据, 处理完后通知外面
两个方法:
	1.处理完后, 往一个新的channel发送处理完的信号, 外面的goroutine去收这个信号
	2.使用WaitGroup
*/

package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(i int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d recive %c\n", i, n)
		w.wg.Done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(i, w)
	return w
}

func channelDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	channelDemo()
}
