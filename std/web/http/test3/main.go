/**
1.在浏览器打开文件 .html, 点击提交浏览器向服务器(可以在action中指定)提交(post)一个表单
2.服务器收到, 打印出提交的键值对
3.接收文件
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func process(wr http.ResponseWriter, r *http.Request) {
	//1.各字段的内容
	//r.ParseForm()
	//fmt.Fprintln(wr, r.Form)
	//fmt.Fprintln(wr, r.PostForm)
	//
	//r.ParseMultipartForm(1024)
	//fmt.Fprintln(wr, r.MultipartForm)

	//2.直接用函数
	fmt.Fprintln(wr, r.FormValue("name"))
	fmt.Fprintln(wr, r.PostFormValue("name"))
}

func processFile(wr http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(wr, string(data))
		}
	}
}

func main() {
	http.HandleFunc("/process", process)
	http.HandleFunc("/file", processFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
