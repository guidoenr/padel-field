package main

import (
	"github.com/guidoenr/padel-field/dbase"
	"github.com/guidoenr/padel-field/turnero"
)

func main() {
	dbase.Init()
	turnero.InitializeTurnos()
}
