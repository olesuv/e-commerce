package tests

import (
	"testing"

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
