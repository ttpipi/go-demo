/**
串联处理器, 也就是拦截器
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "call hello\n")
}

func Debug(f http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(wr, "call log\n")
		f(wr, r)
	}
}

func main() {
	http.HandleFunc("/hello/", Debug(hello))
	log.Fatal(http.ListenAndServe(":1234", nil))
}
