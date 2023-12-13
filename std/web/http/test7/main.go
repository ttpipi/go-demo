/**
模板函数
*/
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var DirPath string = "./web/test7/"

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(rw http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("t1.html").Funcs(funcMap)
	t, _ = t.ParseFiles(DirPath + "t1.html")
	t.Execute(rw, time.Now())
}

func main() {
	http.HandleFunc("/process", process)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
