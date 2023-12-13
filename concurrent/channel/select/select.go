/***
	select调度问题
	1: 收到的数据要给另外的goroutine使用(也就是送给另一个channel), 这个动作什么阻塞
  	2. 收到数据的速度比消耗数据的速度快, 会导致数据丢失

	time channel:  定时, timeout, tick

	nil
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(i int, ch chan int) {
	for n := range ch {
		fmt.Printf("worker %d recive %d\n", i, n)
		time.Sleep(time.Second)
	}
}

func createWorker(i int) chan<- int {
	ch := make(chan int)
	go worker(i, ch)
	return ch
}

func main() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int //nil, 不会被select到, 如果放到for外面, 赋值后就不是nil, 导致不断的收
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Println("bye")
			return
		//类似于linux中的select超时时间
		case <-time.After(time.Millisecond * 800):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("values len : ", len(values))
		}
	}
}
