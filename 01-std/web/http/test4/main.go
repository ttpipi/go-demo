/**
返回响应, ResponWriter的使用:
1.ResponWriter是一个接口, 指向的是引用类型的参数
2.Write
3.Header, 重定向, cookie
4.WriteHeader
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func AboutHeader(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Location", "https://www.baidu.com/")
	//重定向需要将响应码设置为302
	wr.WriteHeader(302)
}

func AboutWrite(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello world!\n"))
}

func AboutWriteHeader(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(501)
	fmt.Fprintln(wr, "No such service, try next door")
}

type Student struct {
	name  string
	score int
}

func AboutJson(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	s := &Student{
		name:  "linzexin",
		score: 100,
	}
	json, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}
	io.WriteString(wr, string(json))
}

func main() {
	http.HandleFunc("/header", AboutHeader)
	http.HandleFunc("/write", AboutWrite)
	http.HandleFunc("/writeHeader", AboutWriteHeader)
	http.HandleFunc("/json", AboutJson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
