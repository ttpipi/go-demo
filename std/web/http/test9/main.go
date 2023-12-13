/**
自定义模板
*/
package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var dirPath string = "./web/test9/"

func process(rw http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	n := rand.Intn(10)
	if n >= 5 {
		t, _ = template.ParseFiles(dirPath+"layout.html", dirPath+"blue.html")
	} else {
		t, _ = template.ParseFiles(dirPath+"layout.html", dirPath+"red.html")
	}

	t.ExecuteTemplate(rw, "layout", "")
}

func processBlock(rw http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	n := rand.Intn(9)
	if n < 3 {
		t, _ = template.ParseFiles(dirPath+"block.html", dirPath+"blue.html")
	} else if n >= 3 && n < 6 {
		t, _ = template.ParseFiles(dirPath+"block.html", dirPath+"red.html")
	} else {
		t, _ = template.ParseFiles(dirPath + "block.html")
	}
	t.ExecuteTemplate(rw, "layout", "")
}

func main() {
	//同一个文件定义多个模板, 同一个模板名字定义在多个文件
	http.HandleFunc("/process", process)

	//块动作
	http.HandleFunc("/process_block", processBlock)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
