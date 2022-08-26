package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

// user roles
const (
	ADMIN  string = "admin"
	NORMAL string = "normal"
)

// Register has the purpose of authenticate the user credentials
func Register(user *models.User) error {

	user.Role = NORMAL
	db := models.InitDB()

	// get all the usernames/email from the database
	exists, err := db.NewSelect().
		Table("users").
		Column("username").
		Where("username = ?", user.Username).
		Exists(context.Background())

	if exists {
		msg := fmt.Sprintf("username '%s' exists \n", user.Username)
		logger.Logerror.Println(msg)
		return errors.New(msg)
	}

	// inserting the new User in the database
	_, err = db.NewInsert().
		Model(user).
		Exec(context.Background())

	if err != nil {
		logger.Logerror.Printf("can't persist in the db: %v", err)
	}

	defer db.Close()
	return err
}

func InitializeUsers() {
	db := models.InitDB()
	rootUser := models.User{
		Username:  "root",
		Password:  "rootoor",
		Role:      ADMIN,
		Phone:     "",
		Email:     "",
		Firstname: "Root",
		Lastname:  "Root",
	}

	_, _ = db.NewInsert().
		Model(&rootUser).
		Exec(context.Background())

	defer db.Close()
}
