package resolvers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/constants"
	"server.go/graph/model"
	"server.go/middleware"
	"server.go/models"
	"server.go/services"
	errors "server.go/utils/errors"
	typesConverters "server.go/utils/types_converters"
)

type OrderResolver struct {
	orderService       *services.OrderService
	orderErrors        errors.OrderErrors
	orderTypeConverter typesConverters.OrderTypesConverter
	userService        *services.UserService
}

func NewOrderResolver() *OrderResolver {
	return &OrderResolver{
		orderService: services.NewOrderService(),
		userService:  services.NewUserService(),
	}
}

func (r *OrderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
	return obj.Id.Hex(), nil
}

func (r *OrderResolver) Currency(ctx context.Context, obj *models.Order) (model.Currency, error) {
	var currency model.Currency
	switch obj.Currency {
	case constants.USD:
		currency = model.CurrencyUsd
	case constants.EUR:
		currency = model.CurrencyEur
	default:
		currency = model.CurrencyUah
	}
	return currency, nil
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

func (r *OrderResolver) Status(ctx context.Context, obj *models.Order) (model.Status, error) {
	var status model.Status
	switch obj.Status {
	case constants.Available:
		status = model.StatusAvailable
	case constants.Buyed:
		status = model.StatusArchived
	default:
		status = model.StatusAvailable
	}
	return status, nil
}

func (r *OrderResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*models.Order, error) {
	if userEmail := middleware.CtxValue(ctx); userEmail == "" {
		return nil, fmt.Errorf("login first")
	}

	err := r.orderErrors.CheckCreateOrderInput(input)
	if err != nil {
		return nil, err
	}

	orderCategories := r.orderTypeConverter.ConvertCategoryTypes(input)
	orderCurrency := r.orderTypeConverter.ConvertCurrencyTypes(input)

	order := &models.Order{
		Id:          primitive.NewObjectID(),
		Title:       *input.Title,
		Description: *input.Description,
		Images:      nil,
		Category:    orderCategories,
		Date:        time.Now(),
		Status:      constants.Available,
		Price:       *input.Price,
		Currency:    orderCurrency,
		AuthorEmail: middleware.CtxValue(ctx),
	}

	order, err = r.orderService.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("server: create order, details: %w", err)
	}

	return order, nil
}

func (r *OrderResolver) BuyOrder(ctx context.Context, orderID string, customerEmail string) (*models.Order, error) {
	if userEmail := middleware.CtxValue(ctx); userEmail == "" {
		return nil, fmt.Errorf("login first")
	}

	_, err := r.userService.GetUserByEmail(customerEmail)
	if err != nil {
		return nil, fmt.Errorf("server: get user, details: %w", err)
	}

	order, err := r.orderService.GetOrderById(orderID)
	if err != nil {
		return nil, fmt.Errorf("server: get order, details: %w", err)
	}
	if order.Status == constants.Buyed {
		return nil, fmt.Errorf("order already buyed")
	}

	order.Status = constants.Buyed
	order.CustomerEmail = customerEmail

	order, err = r.orderService.UpdateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("server: update order, details: %w", err)
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

func (r *OrderResolver) LatestOrders(ctx context.Context) ([]*models.Order, error) {
	orders, err := r.orderService.Last5Orders()
	if err != nil {
		return nil, err
	}

	var orderPointers []*models.Order
	for _, order := range orders {
		orderPointers = append(orderPointers, &order)
	}

	return orderPointers, nil
}

func (r *OrderResolver) SearchOrder(ctx context.Context, userInput string) ([]*models.Order, error) {
	if userInput == "" {
		return nil, nil
	}

	orders, err := r.orderService.SearchByString(userInput)
	if err != nil {
		return nil, err
	}

	var orderPointers []*models.Order
	for _, order := range orders {
		orderPointers = append(orderPointers, &order)
	}

	return orderPointers, nil
}
