package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var numCalcsCreated int32

func createBuffer() any {
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	numWorkers := 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i<numWorkers; i++ {
		go func ()  {
			defer wg.Done()
			buff := bufferPool.Get()
			_ = buff.(*[]byte)
			bufferPool.Put(buff)
		}()
	} 

	wg.Wait()
	fmt.Println("numCalcsCreated=", numCalcsCreated)
}