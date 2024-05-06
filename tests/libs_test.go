package tests

import (
	"context"
	"testing"

	"server.go/configs"
	"server.go/constants"
	"server.go/libs"
)

// TODO: Add user session token generation and validation tests (after docker + redis init)

func TestGenearteJwtToken(t *testing.T) {
	configs.LoadEnv()
	ctx := context.Background()

	token, err := libs.GenearteJwtToken(ctx, constants.EMAIL)
	if err != nil {
		t.Errorf("Failed to generate JWT token: %v", err)
	}
	if token == "" {
		t.Error("Generated JWT token is empty")
	}
}

func TestValidateJwtToken(t *testing.T) {
	configs.LoadEnv()
	ctx := context.Background()

	tokenString, err := libs.GenearteJwtToken(ctx, constants.EMAIL)
	if err != nil {
		t.Errorf(err.Error())
	}

	userEmail, err := libs.ValidateJwtToken(ctx, tokenString)

	if err != nil {
		t.Errorf(err.Error())
	}
	if userEmail == "" {
		t.Error("Failed to validate JWT token")
	}
	if userEmail != constants.EMAIL {
		t.Errorf("Expected email is %s", constants.EMAIL)
	}
}

func TestHashPassword(t *testing.T) {
	hashedPassword := libs.HashPassword(constants.PASSWORD)

	if !libs.VerifyPassword(constants.PASSWORD, hashedPassword) {
		t.Error("Failed to verify hashed password")
	}
}

func TestVerifyPassword(t *testing.T) {
	hashedPassword := libs.HashPassword(constants.PASSWORD)

	if !libs.VerifyPassword(constants.PASSWORD, hashedPassword) {
		t.Error("Failed to verify hashed password")
	}

	invalidPassword := "wrongpassword"
	if libs.VerifyPassword(invalidPassword, hashedPassword) {
		t.Error("Incorrectly verified invalid password")
	}
}
