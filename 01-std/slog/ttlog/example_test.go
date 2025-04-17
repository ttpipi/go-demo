package ttlog_test

import (
	"context"
	"demo/01-std/slog/ttlog"
)

func ExampleTTLog() {
	ttlog.Log(context.Background(), 1, "greeting")

	// Output:
	// greeting
}
