package controllers

import (
	"context"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"net/http"
	"strconv"
	"time"
)

const (
	SecretKey = "secret"
)

// Register has the purpose of authenticate the user credentials
func Register(newUser *models.User) error {
	db := models.InitDB()
	newUser.Role = models.NORMAL

	// get all the usernames from the database
	exists, err := db.NewSelect().
		Table("users").
		Column("username").
		Where("username = ?", newUser.Username).
		Exists(context.Background())

	if exists {
		msg := fmt.Sprintf("username '%s' exists \n", newUser.Username)
		logger.Logerror.Println(msg)
		return errors.New(msg)
	}

	// get all the emails from the database
	exists, err = db.NewSelect().
		Table("users").
		Column("email").
		Where("email = ?", newUser.Email).
		Exists(context.Background())

	if exists {
		msg := fmt.Sprintf("email '%s' exists \n", newUser.Email)
		logger.Logerror.Println(msg)
		return errors.New(msg)
	}

	// inserting the new User in the database
	_, err = db.NewInsert().
		Model(newUser).
		Exec(context.Background())

	if err != nil {
		logger.Logerror.Printf("can't persist in the db: %v", err)
	}
	logger.Loginfo.Printf("persisted user {'%s'-%s} into the database", newUser.Username, newUser.Email)

	defer db.Close()
	return err
}

// Login has the purpose of authenticate the user credentials
func Login(user *models.User) (*http.Cookie, error) {
	db := models.InitDB()

	// find if the username or email exists
	exists, err := db.NewSelect().
		Table("users").
		Column("username").
		Where("username = ?", user.Username).
		Exists(context.Background())

	// if the user does not exist
	if !exists {
		msg := fmt.Sprintf("username '%s' does not exist \n", user.Username)
		logger.Logerror.Println(msg)
		return &http.Cookie{}, errors.New(msg)
	}

	var userFound models.User
	err = db.NewSelect().
		Model(&userFound).
		Where("username = ?", user.Username).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("getting user: %v", err)
		logger.Logerror.Println(msg)
		return &http.Cookie{}, errors.New(msg)
	}

	if user.Password != userFound.Password {
		msg := fmt.Sprintf("wrong password")
		logger.Logerror.Println(msg)
		return &http.Cookie{}, errors.New(msg)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.FormatInt(user.ID, 10),
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // 3 hours
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		msg := fmt.Sprintf("could not log in")
		logger.Logerror.Println(msg)
		return &http.Cookie{}, errors.New(msg)
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 3), // 3 hours
		HttpOnly: true,
	}

	defer db.Close()
	return &cookie, err
}
