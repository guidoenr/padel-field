package main

import (
	"flag"
	"fmt"
	"github.com/guidoenr/padel-field/api"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

/*
	author: @github.com/guidoenr
	repo: github.com/guidoenr/padelfield
*/

// resetDbAndCreateSchemas restart the entire DB and create all the schemas
func resetDbAndCreateSchemas() {
	db := models.InitDB()

	// restarting the DB
	err := models.RestartDb(db)
	if err != nil {
		logger.Logerror.Println(err)
	}

	// initializing the turnos
	err = InitializeTurnos()
	if err != nil {
		logger.Logerror.Println(err)
	}

	// initializing the admin user
	err = InitializeUsers()
	if err != nil {
		logger.Logerror.Println(err)
	}

	defer logger.Loginfo.Println("Listen and Serving..")
	defer db.Close()
}

func startServer() {
	api.ListenAndServe() // gin gonic server on 8080
	defer logger.Loginfo.Println("Listen and Serving..")
}

func main() {
	defer logger.Loginfo.Println("starting app..")
	resetDb := flag.Bool("r", false, "[WARNING] this restart the ENTIRE db")
	startSv := flag.Bool("s", false, "this flag start the gin-gonic api")
	flag.Parse()

	if *resetDb {
		fmt.Println("restarting db..")
		resetDbAndCreateSchemas()
	}
	if *startSv {
		fmt.Println("starting server..")
		startServer()
	}

}
