package adapter

type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct{}

func (a AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}
