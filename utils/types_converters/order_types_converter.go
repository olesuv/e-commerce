package utils

import (
	"server.go/constants"
	"server.go/graph/model"
)

type OrderTypesConverter struct{}

func (tc *OrderTypesConverter) ConvertCategoryTypes(input model.CreateOrderInput) constants.OrderCategories {
	orderCategories := make(constants.OrderCategories, 0)
	if input.Category == nil || len(input.Category) == 0 {
		orderCategories = append(orderCategories, constants.Other)
	}
	if len(input.Category) >= 1 {
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
	return orderCategories
}

func (tc *OrderTypesConverter) ConvertCurrencyTypes(input model.CreateOrderInput) constants.OrderCurrency {
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
	return orderCurrency
}
