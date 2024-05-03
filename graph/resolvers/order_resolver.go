package resolvers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/constants"
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

func (r *OrderResolver) Category(ctx context.Context, obj *models.Order) ([]int, error) {
	categories := make([]int, len(obj.Category))
	for i, category := range obj.Category {
		categories[i] = int(category)
	}
	return categories, nil
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

	if input.Category == nil || len(input.Category) == 0 {
		otherCategory := constants.OrderCategory(constants.Other)
		input.Category = make([]*int, 0)
		input.Category = append(input.Category, (*int)(&otherCategory))
	}
	orderCategories := make(constants.OrderCategories, 0)
	for _, category := range input.Category {
		if category == nil || *category < int(constants.OrderCategory(constants.Electronics)) || *category > int(constants.OrderCategory(constants.Other)) {
			return nil, fmt.Errorf("category is invalid")
		}

		orderCategories = append(orderCategories, constants.OrderCategory(*category))
	}

	if input.Images == nil {
		return nil, fmt.Errorf("images are required")
	}

	if input.Price == nil {
		return nil, fmt.Errorf("price is required")
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
		Category:    orderCategories,
		Date:        time.Now(),
		Status:      constants.Available,
		Price:       *input.Price,
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
