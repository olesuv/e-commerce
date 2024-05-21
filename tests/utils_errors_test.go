package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"server.go/graph/model"
	errors "server.go/utils/errors"
)

func TestCheckCreateUserInput(t *testing.T) {
	ue := errors.NewUserErrors()

	// Test case 1: Empty email
	input1 := model.CreateUserInput{
		Email:    nil,
		Password: strPtr("password123"),
	}
	err1 := ue.CheckCreateUserInput(input1)
	if err1 == nil || err1.Error() != "email is required" {
		t.Errorf("Expected 'email is required' error, got: %v", err1)
	}

	// Test case 2: Empty password
	input2 := model.CreateUserInput{
		Email:    strPtr("test@example.com"),
		Password: nil,
	}
	err2 := ue.CheckCreateUserInput(input2)
	if err2 == nil || err2.Error() != "password is required" {
		t.Errorf("Expected 'password is required' error, got: %v", err2)
	}

	// Test case 3: Invalid email address
	input3 := model.CreateUserInput{
		Email:    strPtr("invalid_email"),
		Password: strPtr("password123"),
	}
	err3 := ue.CheckCreateUserInput(input3)
	if err3 == nil || err3.Error() != "invalid email address" {
		t.Errorf("Expected 'invalid email address' error, got: %v", err3)
	}

	// Test case 4: Existing user with the same email
	input4 := model.CreateUserInput{
		Email:    strPtr("test0@test.com"),
		Password: strPtr("password123"),
	}
	err4 := ue.CheckCreateUserInput(input4)
	if err4 == nil || err4.Error() != "user with this email already exists" {
		t.Errorf("Expected 'user with this email already exists' error, got: %v", err4)
	}

	// Test case 5: Valid input
	input5 := model.CreateUserInput{
		Email:    strPtr("new_user@example.com"),
		Password: strPtr("password123"),
	}
	err5 := ue.CheckCreateUserInput(input5)
	if err5 != nil {
		t.Errorf("Expected no error, got: %v", err5)
	}
}

func TestUserErrors_CheckLoginUserInput(t *testing.T) {
	// Create a new instance of UserErrors
	ue := errors.NewUserErrors()

	// Define the input for the method
	input := model.LoginUserInput{
		Email:    strPtr("test0@test.com"),
		Password: strPtr("securepassword"),
	}

	// Call the method
	_, err := ue.CheckLoginUserInput(input)

	// Check if the error is as expected
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestSetName(t *testing.T) {
	ue := errors.NewUserErrors()

	// Test case 1: input.Name is nil
	input1 := &model.CreateUserInput{}
	expected1 := "New User"
	result1 := ue.SetName(input1)
	if result1 != expected1 {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected1, result1)
	}

	// Test case 2: input.Name is an empty string
	input2 := &model.CreateUserInput{Name: strPtr("")}
	expected2 := "New User"
	result2 := ue.SetName(input2)
	if result2 != expected2 {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected2, result2)
	}

	// Test case 3: input.Name is not empty
	input3 := &model.CreateUserInput{Name: strPtr("John Doe")}
	expected3 := "John Doe"
	result3 := ue.SetName(input3)
	if result3 != expected3 {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected3, result3)
	}
}

func TestCheckCreateOrderInput(t *testing.T) {
	orderErrors := &errors.OrderErrors{}

	// Test case: valid input
	validInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("Sample Description"),
		Price:       floatPtr(10.99),
	}
	err := orderErrors.CheckCreateOrderInput(validInput)
	assert.NoError(t, err, "Expected no error for valid input")

	// Test case: empty title
	emptyTitleInput := model.CreateOrderInput{
		Title:       nil,
		Description: strPtr("Sample Description"),
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(emptyTitleInput)
	assert.Error(t, err, "Expected error for empty title")

	// Test case: title length less than 3 characters
	shortTitleInput := model.CreateOrderInput{
		Title:       strPtr("Ti"),
		Description: strPtr("Sample Description"),
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(shortTitleInput)
	assert.Error(t, err, "Expected error for short title")

	// Test case: title length greater than 100 characters
	longTitleInput := model.CreateOrderInput{
		Title:       strPtr("This is a very long title that exceeds the maximum length of 100 characters. This is a very long title that exceeds the maximum length of 100 characters. This is a very long title that exceeds the maximum length of 100 characters."),
		Description: strPtr("Sample Description"),
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(longTitleInput)
	assert.Error(t, err, "Expected error for long title")

	// Test case: empty description
	emptyDescriptionInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: nil,
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(emptyDescriptionInput)
	assert.Error(t, err, "Expected error for empty description")

	// Test case: description length less than 3 characters
	shortDescriptionInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("De"),
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(shortDescriptionInput)
	assert.Error(t, err, "Expected error for short description")

	// Test case: description length greater than 500 characters
	longDescriptionInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum"),
		Price:       floatPtr(10.99),
	}
	err = orderErrors.CheckCreateOrderInput(longDescriptionInput)
	assert.Error(t, err, "Expected error for long description")

	// Test case: price is nil
	nilPriceInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("Sample Description"),
		Price:       nil,
	}
	err = orderErrors.CheckCreateOrderInput(nilPriceInput)
	assert.Error(t, err, "Expected error for nil price")

	// Test case: price less than 0
	negativePriceInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("Sample Description"),
		Price:       floatPtr(-10.99),
	}
	err = orderErrors.CheckCreateOrderInput(negativePriceInput)
	assert.Error(t, err, "Expected error for negative price")

	// Test case: price greater than 1000000
	largePriceInput := model.CreateOrderInput{
		Title:       strPtr("Sample Title"),
		Description: strPtr("Sample Description"),
		Price:       floatPtr(1000001),
	}
	err = orderErrors.CheckCreateOrderInput(largePriceInput)
	assert.Error(t, err, "Expected error for large price")
}
