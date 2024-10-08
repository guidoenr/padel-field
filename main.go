package main

import (
	"flag"
	"github.com/guidoenr/padel-field/api"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models/psdb"
)

func main() {
	logger.Loginfo.Println("starting app..")
	resetDb := flag.Bool("r", false, "[WARNING] this restart the ENTIRE db")
	startSv := flag.Bool("s", false, "this flag start the gin-gonic api")
	flag.Parse()

	if *resetDb {
		logger.Logwarning.Println("RESTARTING DB")
		resetDbAndCreateSchemas()
	}
	if *startSv {
		logger.Loginfo.Println("starting server..")
		startServer()
	}
}

// resetDbAndCreateSchemas restart the entire DB and create all the schemas
func resetDbAndCreateSchemas() {
	db := psdb.InitDB()

	// restarting the DB
	err := psdb.RestartDb(db)
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

	defer db.Close()
}

func startServer() {
	logger.Loginfo.Println("listen and serving on http://localhost:8080")
	logger.Loginfo.Println("si llegaste hasta aca marcos felicidades.. y bienvenido a la magia de docker")
	api.ListenAndServe() // gin gonic server on 8080
}
