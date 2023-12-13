package main

import (
	"context"
	"fmt"
)

type B struct {
	h Handler
}

func NewB() HandlerMiddleware {
	return func(handler Handler) Handler {
		return B{h: handler}
	}
}

func (b B) Handle(ctx context.Context) error {
	fmt.Println("B")
	return b.h.Handle(ctx)
}
