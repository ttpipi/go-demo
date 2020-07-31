package adapter

type Target interface {
	Request() string
}

type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

func (a adapter) Request() string {
	return a.SpecificRequest()
}
