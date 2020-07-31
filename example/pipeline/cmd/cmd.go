package main

import (
	"bufio"
	"fmt"
	"golang/example/pipeline/pipeline"
	"os"
)

func main() {
	// test1()

	const fileName = "small.in"
	const n = 64
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func test1() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 6, 4, 5, 9)),
		pipeline.InMemSort(pipeline.ArraySource()))

	for v := range p {
		fmt.Println(v)
	}
}
