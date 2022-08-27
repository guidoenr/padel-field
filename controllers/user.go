package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

// Register has the purpose of authenticate the user credentials
func Register(user *models.User) error {

	user.Role = models.NORMAL
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
