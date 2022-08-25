package main

import (
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

/*
	author: @github.com/guidoenr
	repo: github.com/guidoenr/padelfield
*/

func main() {

	//ListenAndServe() // Listen and Serve on 8080
	db := models.InitDB()
	//
	err := models.RestartDb(db)
	if err != nil {
		logger.Logerror.Println(err)
	}

	err2 := models.CreateSchemas(db)
	if err2 != nil {
		logger.Logerror.Println(err2)
	}

	defer db.Close()

	//UpdateTurnos()
	//InitializeTurnos()
	//
	//dummyUser := models.User{
	//	Username:  "guidoenr",
	//	Password:  "aleatoryPassword",
	//	Phone:     "2325684951",
	//	Email:     "guidoenr4@gmail.com",
	//	Firstname: "Guido",
	//	Lastname:  "Enrique",
	//}
	//
	//controllers.Register(&dummyUser)
}
