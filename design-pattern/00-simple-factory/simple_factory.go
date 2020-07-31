package simple_factory

import "fmt"

type Api interface {
	Say(name string)
}

//有新类时, 需要修改工厂函数
func NewApi(t int) Api {
	switch t {
	case 1:
		return &SayHi{}
	case 2:
		return &SayHello{}
	}
	return nil
}

//SayHi
type SayHi struct{}

func (s *SayHi) Say(name string) {
	fmt.Println("Hi, " + name)
}

//SayHello
type SayHello struct{}

func (s *SayHello) Say(name string) {
	fmt.Println("Hello, " + name)
}
