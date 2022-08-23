package main

import (
	"github.com/guidoenr/padel-field/db"
	"github.com/guidoenr/padel-field/turnero"
)

func main() {
	//initialize()
	db.Init()
	turnero.InitializeTurnos()
}
