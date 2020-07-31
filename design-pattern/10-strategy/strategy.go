package strategy

import "fmt"

type PaymentContext struct {
	Name    string
	CardID  string
	Money   int
	payment PaymentStrategy
}

func NewPaymentContext(name string, cardID string, money int, payment PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Name:    name,
		CardID:  cardID,
		Money:   money,
		payment: payment,
	}
}

func (p *PaymentContext) Pay() {
	p.payment.Pay(p)
}

//策略类
type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Cash struct{}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct{}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}
