package proxy

type Subject interface {
	Do() string
}

//真实的对象
type RealSubject struct{}

func (r RealSubject) Do() string {
	return "real"
}

//代理的对象
type ProxySubject struct {
	real RealSubject
}

func (p ProxySubject) Do() string {
	var res string
	res += "pre:"

	res += p.real.Do()

	res += ":after"
	return res
}
