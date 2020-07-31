package strategy

import "testing"

func TestStrategy(t *testing.T) {
	ctx := NewPaymentContext("lin", "", 100, &Cash{})
	ctx.Pay()

	ctx = NewPaymentContext("zexin", "123456", 100, &Bank{})
	ctx.Pay()
}
