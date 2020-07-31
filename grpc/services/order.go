package services

import (
	context "context"
	"fmt"
)

type OrderService struct{}

func (*OrderService) NewOrder(_ context.Context, order *OrderRequest) (*OrderResponse, error) {
	fmt.Println(order.Order)
	return &OrderResponse{
		Status:  "ok",
		Message: "success",
	}, nil
}
