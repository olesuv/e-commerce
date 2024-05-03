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

func (r *OrderResolver) Category(ctx context.Context, obj *models.Order) ([]model.Category, error) {
	var categories []model.Category
	for _, category := range obj.Category {
		switch category {
		case constants.Electronics:
			categories = append(categories, model.CategoryElectronics)
		case constants.Fashion:
			categories = append(categories, model.CategoryFashion)
		case constants.Home:
			categories = append(categories, model.CategoryHome)
		case constants.Sports:
			categories = append(categories, model.CategorySports)
		case constants.Books:
			categories = append(categories, model.CategoryBooks)
		case constants.Automotive:
			categories = append(categories, model.CategoryAutomotive)
		case constants.Other:
			categories = append(categories, model.CategoryOther)
		}
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

	orderCategories := make(constants.OrderCategories, len(input.Category))
	if input.Category == nil || len(input.Category) == 0 {
		orderCategories = append(orderCategories, constants.Other)
	}
	if len(input.Category) > 1 {
		for _, category := range input.Category {
			switch *category {
			case model.CategoryElectronics:
				orderCategories = append(orderCategories, constants.Electronics)
			case model.CategoryFashion:
				orderCategories = append(orderCategories, constants.Fashion)
			case model.CategoryHome:
				orderCategories = append(orderCategories, constants.Home)
			case model.CategorySports:
				orderCategories = append(orderCategories, constants.Sports)
			case model.CategoryBooks:
				orderCategories = append(orderCategories, constants.Books)
			case model.CategoryAutomotive:
				orderCategories = append(orderCategories, constants.Automotive)
			case model.CategoryOther:
				orderCategories = append(orderCategories, constants.Other)
			default:
				orderCategories = append(orderCategories, constants.Other)
			}
		}
	}

	if input.Images == nil {
		return nil, fmt.Errorf("images are required")
	}

	if input.Price == nil {
		return nil, fmt.Errorf("price is required")
	}

	var orderCurrency constants.OrderCurrency = constants.UAH
	if input.Currency != nil {
		switch *input.Currency {
		case model.CurrencyUsd:
			orderCurrency = constants.USD
		case model.CurrencyEur:
			orderCurrency = constants.EUR
		default:
			orderCurrency = constants.UAH
		}
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
		Currency:    orderCurrency,
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
