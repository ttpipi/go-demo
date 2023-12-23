/*
	往关闭的channel写数据会导致panic
	往关闭的channel读数据则读到零值, 如果是指针, 会读到nil, 使用nil也会panic
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan *int, 3)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan *int) {
		x := 1
		ch <- &x
		x = 2
		ch <- &x
		x = 3
		ch <- &x
		close(ch)

		//这里会panic
		x = 4
		ch <- &x

		fmt.Println("goroutine1 exit")
		wg.Done()
	}(ch)

	go func(ch chan *int) {
		time.Sleep(time.Second * 3)
		for {
			select {
			case <-time.After(time.Second):
				//
			case x := <-ch:
				//操作nil指针引起panic
				n := *x
				fmt.Println(n)
			}
		}
		wg.Done()
	}(ch)

	wg.Wait()
	fmt.Println("app exit")
}
