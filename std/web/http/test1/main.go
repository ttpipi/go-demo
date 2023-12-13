package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
启动一个简单的web服务器
*/

func test1() {
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(wr, "hello world")
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}

func test2() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	}))
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func main() {
	//test1()
	test2()
}
