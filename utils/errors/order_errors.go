package utils

import (
	"fmt"
	"unicode/utf8"

	"server.go/graph/model"
)

type OrderErrors struct{}

func (o *OrderErrors) CheckCreateOrderInput(input model.CreateOrderInput) error {
	if input.Title == nil || *input.Title == "" {
		return fmt.Errorf("title is required")
	}
	if utf8.RuneCountInString(*input.Title) < 3 {
		return fmt.Errorf("minimum length of title is 3 characters")
	}
	if utf8.RuneCountInString(*input.Title) > 100 {
		return fmt.Errorf("maximum length of title is 100 characters")
	}

	if input.Description == nil || *input.Description == "" {
		return fmt.Errorf("description is required")
	}
	if utf8.RuneCountInString(*input.Description) < 3 {
		return fmt.Errorf("minimum length of descriprion is 3 characters")
	}
	if utf8.RuneCountInString(*input.Description) > 500 {
		return fmt.Errorf("maximum length of title is 500 characters")
	}

	if input.Price == nil {
		return fmt.Errorf("price is required")
	}
	if *input.Price < float64(0) {
		return fmt.Errorf("price must be greater than 0")
	}
	if *input.Price > float64(1000000) {
		return fmt.Errorf("this platform is for 'normal' people, not for millionaires")
	}

	return nil
}
