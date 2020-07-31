package facade

//如果新加入子系统C, 该外观对象也要修改, 违反开闭原则, 可以引入抽象外观类
type System struct {
	a SystemA
	b SystemB
}

func NewSystem() System {
	return System{
		a: SystemA{},
		b: SystemB{},
	}
}

func (s System) Run() {
	s.a.Run()
	s.b.Run()
}
