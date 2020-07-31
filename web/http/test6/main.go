/**
模板和模板动作
*/
package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var DirPath string = "./web/test6/"

func process(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(DirPath + "tmpl1.html")
	t.Execute(rw, "hello world!\n")
}

func processIf(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(DirPath + "tmpl2.html")
	rand.Seed(time.Now().Unix())
	t.Execute(rw, rand.Intn(10) > 5)
}

func processRange(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(DirPath + "tmpl3.html")
	ar := []string{"ttlin", "linzexin", "shashou"}
	t.Execute(rw, ar)
}

func processWith(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(DirPath + "tmpl4.html")
	t.Execute(rw, "hello")
}

func processInclude(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(DirPath+"t5.html", DirPath+"t6.html")
	t.Execute(rw, "hello world!")
}

func main() {
	http.HandleFunc("/process", process)
	//条件动作
	http.HandleFunc("/process_if", processIf)
	//迭代动作
	http.HandleFunc("/process_range", processRange)
	//设置动作
	http.HandleFunc("/process_with", processWith)
	//包含动作
	http.HandleFunc("/process_include", processInclude)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
