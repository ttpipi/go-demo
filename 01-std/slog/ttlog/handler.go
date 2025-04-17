package ttlog

import "context"

type Handler interface {
	Handle(context.Context, any) error
}

type defaultHandler struct {
}

func newDefaultHandler() *defaultHandler {
	return &defaultHandler{}
}

func (d *defaultHandler) Handle(p0 context.Context, p1 any) error {
	panic("TODO: Implement")
}
