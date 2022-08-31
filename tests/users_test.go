package main

import (
	"fmt"
	"github.com/guidoenr/padel-field/api/controllers"
	"github.com/guidoenr/padel-field/models"
	"testing"
)

/*
TESTING payload fields
*/
func TestUsernameExists(t *testing.T) {
	johnDoe := models.User{
		Username:  "johndoe123",
		Password:  "easy_password123[';",
		Phone:     "51250120",
		Email:     "johndoe@outlook.com",
		Firstname: "John",
		Lastname:  "Doe",
	}

	// first register
	_ = controllers.Register(&johnDoe)

	// second register
	err := controllers.Register(&johnDoe)

	if err != nil {
		fmt.Println("error")
	}
}
