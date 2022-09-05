package controllers

import (
	"context"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/guidoenr/padel-field/api/errs"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"github.com/guidoenr/padel-field/models/psdb"
	"net/http"
	"strconv"
	"time"
)

const (
	SecretKey = "secret"
)

// Register has the purpose of authenticate the user credentials
func Register(newUser *models.User) errs.RequestError {
	db := psdb.InitDB()
	newUser.Role = models.NORMAL

	// get all the usernames from the database
	exists, err := db.NewSelect().
		Table("users").
		Column("username").
		Where("username = ?", newUser.Username).
		Exists(context.Background())

	if exists {
		return errs.ThrowError(errors.New("select error"), "username exists", 1)
	}

	// get all the emails from the database
	exists, err = db.NewSelect().
		Table("users").
		Column("email").
		Where("email = ?", newUser.Email).
		Exists(context.Background())

	if exists {
		return errs.ThrowError(err, "email exists", 2)
	}

	// inserting the new User in the database
	_, err = db.NewInsert().
		Model(newUser).
		Exec(context.Background())

	if err != nil {
		return errs.ThrowError(err, "cannot persist into the db", -1)
	}

	logger.Loginfo.Printf("persisted user '%s' into the database", newUser.Username)

	defer db.Close()
	return errs.ThrowEmptyError()
}

// Login has the purpose of authenticate the user credentials
func Login(user *models.User) (*http.Cookie, errs.RequestError) {
	db := psdb.InitDB()

	// find if the username or email exists
	exists, err := db.NewSelect().
		Table("users").
		Column("username").
		Where("username = ?", user.Username).
		Exists(context.Background())

	// if the user does not exist
	if !exists {
		return &http.Cookie{}, errs.ThrowError(err, "username does not exist", 3)
	}

	// the username exists
	var userFound models.User
	err = db.NewSelect().
		Model(&userFound).
		Where("username = ?", user.Username).
		Scan(context.Background())

	if err != nil {
		return &http.Cookie{}, errs.ThrowError(err, "cannot get user from db", -1)
	}

	if user.Password != userFound.Password {
		return &http.Cookie{}, errs.ThrowError(errors.New("password does not match"), "password does not match", 4)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.FormatInt(userFound.ID, 10),
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // 3 hours
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return &http.Cookie{}, errs.ThrowError(err, "could not get token", -1)
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 3), // 3 hours
		HttpOnly: true,
	}

	defer db.Close()
	return &cookie, errs.ThrowEmptyError()
}

func GetUserById(id string) (models.User, errs.RequestError) {
	var user models.User
	db := psdb.InitDB()

	err := db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(context.Background())

	if err != nil {
		return user, errs.ThrowError(err, "can not get user by id", -1)
	}

	defer db.Close()

	return user, errs.ThrowEmptyError()
}
