package main

import (
	"github.com/guidoenr/padel-field/api"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

/*
	author: @github.com/guidoenr
	repo: github.com/guidoenr/padelfield
*/

// cleanAll restart the entire DB and create all the schemas
func cleanAll() {
	db := models.InitDB()

	err := models.RestartDb(db)
	if err != nil {
		logger.Logerror.Println(err)
	}

	InitializeTurnos()
	InitializeUsers()

	defer db.Close()
}

func main() {
	cleanAll()
	api.ListenAndServe() // gin gonic server on 8080
	UpdateTurnos()
	//dummyUser := models.User{
	//	Username:  "johndoe123",
	//	Password:  "easy_password123[';",
	//	Phone:     "51250120",
	//	Email:     "johndoe@outlook.com",
	//	Firstname: "John",
	//	Lastname:  "Doe",
	//}
	//err := controllers.Register(&dummyUser)
	//if err != nil {
	//	logger.Logerror.Printf("register: %v", err)
	//}
}
