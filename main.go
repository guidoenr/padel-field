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

}
