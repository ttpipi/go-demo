package main

import (
	"context"
	"fmt"
)

type A struct {
	h Handler
}

func NewA() HandlerMiddleware {
	return func(handler Handler) Handler {
		return A{h: handler}
	}
}

func (a A) Handle(ctx context.Context) error {
	fmt.Println("A")
	return a.h.Handle(ctx)
}
