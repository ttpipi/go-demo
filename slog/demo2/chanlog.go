package main

import (
	"bytes"
	"context"
	"log/slog"
)

type ChanHandler struct {
	slog.Handler
	ch  chan []byte
	buf *bytes.Buffer
}

func (h *ChanHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *ChanHandler) Handle(ctx context.Context, r slog.Record) error {
	if err := h.Handler.Handle(ctx, r); err != nil {
		return err
	}
	var nb = make([]byte, h.buf.Len())
	copy(nb, h.buf.Bytes())
	h.ch <- nb
	h.buf.Reset()
	return nil
}

func (h *ChanHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return &ChanHandler{
		Handler: h.Handler.WithAttrs(as),
		ch:      h.ch,
		buf:     h.buf,
	}
}

func (h *ChanHandler) WithGroup(name string) slog.Handler {
	return &ChanHandler{
		Handler: h.Handler.WithGroup(name),
		ch:      h.ch,
		buf:     h.buf,
	}
}

func NewChanHandler(ch chan []byte) *ChanHandler {
	h := &ChanHandler{
		ch:  ch,
		buf: bytes.NewBuffer(nil),
	}

	h.Handler = slog.NewJSONHandler(h.buf, nil)
	return h
}
