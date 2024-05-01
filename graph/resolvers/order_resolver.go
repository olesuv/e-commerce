package resolvers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/graph/model"
	"server.go/libs"
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

func (r *OrderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
	return obj.Id.Hex(), nil
}

func (r *OrderResolver) Category(ctx context.Context, obj *models.Order) (int, error) {
	return int(obj.Category), nil
}

func (r *OrderResolver) Status(ctx context.Context, obj *models.Order) (int, error) {
	return int(obj.Status), nil
}

func (r *OrderResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*models.Order, error) {
	if input.Title == nil || *input.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if len(*input.Title) <= 3 {
		return nil, fmt.Errorf("minimum length of title is 3 characters")
	}
	if len(*input.Title) >= 100 {
		return nil, fmt.Errorf("maximum length of title is 100 characters")
	}

	if input.Description == nil || *input.Description == "" {
		return nil, fmt.Errorf("description is required")
	}
	if len(*input.Description) <= 3 {
		return nil, fmt.Errorf("minimum length of description is 10 characters")
	}
	if len(*input.Description) >= 1000 {
		return nil, fmt.Errorf("maximum length of description is 1000 characters")
	}

	if input.Category == nil {
		return nil, fmt.Errorf("category is required")
	}

	if input.Images == nil {
		return nil, fmt.Errorf("images are required")
	}

	compressedImgs := []primitive.Binary{}
	for _, img := range input.Images {
		compressedImg, err := libs.CompressImage(*img)
		if err != nil {
			return nil, err
		}

		binImg := primitive.Binary{
			Data: compressedImg,
		}

		compressedImgs = append(compressedImgs, binImg)
	}

	order := &models.Order{
		Id:          primitive.NewObjectID(),
		Title:       *input.Title,
		Description: *input.Description,
		Images:      compressedImgs,
		Category:    models.OrderCategory(*input.Category),
		Date:        time.Now(),
		Status:      models.Available,
	}

	order, err := r.orderService.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("server: create order, details: %w", err)
	}

	return order, nil
}

func (r *OrderResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	orders, err := r.orderService.GetOrders()
	if err != nil {
		return nil, fmt.Errorf("server: get orders, details: %w", err)
	}

	var orderPointers []*models.Order
	for _, order := range orders {
		orderPointers = append(orderPointers, &order)
	}

	return orderPointers, nil
}

func (r *OrderResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	order, err := r.orderService.GetOrderById(id)
	if err != nil {
		return nil, fmt.Errorf("server: get order, details: %w", err)
	}

	return order, nil
}
