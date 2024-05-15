package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hashedPassword)
}

func unhashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyPassword(userPassword string, userHash string) (bool, error) {
	comapre := unhashPassword(userPassword, userHash)

	if !comapre {
		return false, fmt.Errorf("invalid password")
	}

	return true, nil
}
