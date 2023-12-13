package main

import (
	"fmt"

	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var count int32

func main() {
	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		// sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// res, _ := singleflightGetArticle(sg, 1)
			res, _ := getArticle(1)
			if res != "article:1" {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s\n", n, time.Since(now))
}

func getArticle(id int32) (article string, err error) {
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	return fmt.Sprintf("article:%d", id), nil
}

func singleflightGetArticle(sg *singleflight.Group, id int32) (article string, err error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})
	return v.(string), err
}
