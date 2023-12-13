package main

import (
	"context"
	"fmt"
)

type Handler interface {
	Handle(ctx context.Context) error
}

type HandlerMiddleware func(handler Handler) Handler

func HandleWithMiddleware(handler Handler, hmw ...HandlerMiddleware) Handler {
	h := handler
	for i := len(hmw) - 1; i >= 0; i-- {
		h = hmw[i](h)
	}
	return h
}

// -------------------------HandlerFunc--------------------------//
type HandlerFunc func(ctx context.Context) error

func (f HandlerFunc) Handle(ctx context.Context) error {
	return f(ctx)
}

func NewHandler() HandlerFunc {
	return func(ctx context.Context) error {
		fmt.Println("HHH")
		return nil
	}
}
