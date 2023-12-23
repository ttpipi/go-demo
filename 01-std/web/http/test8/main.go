/**
上下文感知, XSS攻击
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

var dirPath string = "./web/test8/"

func process(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(dirPath + "t2.html")
	//golang自动可防御XSS攻击
	//t.Execute(rw, r.FormValue("comment"))

	//不防御XSS攻击
	rw.Header().Set("X-XSS-Protection", "0")
	t.Execute(rw, template.HTML(r.FormValue("comment")))
}

func form(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(dirPath + "t1.html")
	t.Execute(rw, nil)
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/process", process)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
