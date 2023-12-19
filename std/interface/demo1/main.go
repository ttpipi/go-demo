package main

import "fmt"

type Writer interface {
	Write()
}

type File interface {
	Open()
	Close()
}

type file struct{}

func (f file) Open()  { fmt.Println("open") }
func (f file) Close() { fmt.Println("close") }
func (f file) Write() { fmt.Println("write") }

func Bar(f File) {
	f.Open()
}

func main() {
	var w Writer = new(file)
	//必须判断, ok了才能用
	if f, ok := w.(File); ok {
		Bar(f)
	}
}
