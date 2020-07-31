package facade

import "fmt"

//子系统A
type SystemA struct{}

func (s SystemA) Run() {
	fmt.Println("system A running")
}

//子系统B
type SystemB struct{}

func (s SystemB) Run() {
	fmt.Println("system B running")
}
