package resolvers

import (
	"context"
	"fmt"

	"server.go/graph/model"
	"server.go/models"
	"server.go/services"
)

type OrderResolver struct {
	orderService *services.OrderService
}

func NewOrderResolver() *OrderResolver {
	return &OrderResolver{
		orderService: services.NewOrderService(),
	}
}

func (r *OrderResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*models.Order, error) {
	return nil, fmt.Errorf("not implemented: CreateOrder - createOrder")
}
