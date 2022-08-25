package controllers

import (
	"context"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

func Register(user *models.User) error {

	var usernames map[string]interface{}
	var err error
	db := models.InitDB()

	_ = db.NewSelect().
		Model(&usernames).
		Table("users").
		Column("id", "email").
		Scan(context.Background())

	_, exists := usernames[user.Username]
	if exists {
		logger.Logerror.Printf("the username '%s' already exists: %v", user.Username, err)
	} else {
		_, err = db.NewInsert().
			Model(user).
			Exec(context.Background())

		if err != nil {
			logger.Logerror.Printf("can't persist in the db: %v", err)
		}
	}

	defer db.Close()
	return err
}
