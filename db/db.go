package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/guidoenr/padel-field/models"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

type EnvConfig struct {
	Username string
	Password string
	Database string
	Port     string
}

// Init only connects to the existing DB
func Init() {

	cfg := loadEnv()
	db := pg.Connect(&pg.Options{
		Database: cfg.Database,
		User:     cfg.Username,
		Password: cfg.Password,
	})

	err := db.Ping(context.Background())
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer db.Close()
	err = createSchema(db)
	if err != nil {
		fmt.Printf("error creating schema: %v", err)
	}

	user1 := models.User{
		Id:        0,
		Username:  "guido",
		Email:     "guido@email.com",
		Password:  "random123",
		Firstname: "Guido",
		Lastname:  "Enrique",
	}

	_, err = db.Model(&user1).Insert()
	if err != nil {
		fmt.Printf("error inserting in the database: %v", err)
	}

	// Select all users.
	var users []models.User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	modelsMap := []interface{}{
		(*models.User)(nil),
		(*models.Turno)(nil),
	}

	// create the tables
	for _, model := range modelsMap {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// loadEnv load the environment variables from `local.env` // TODO, change later to heroku maybe?
func loadEnv() EnvConfig {
	err := godotenv.Load("db/local.env")
	if err != nil {
		log.Fatalf("error reading 'local.env' file err: %v", err)
	}

	cfg := EnvConfig{
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
	}
	return cfg
}
