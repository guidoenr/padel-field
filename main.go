package main

import (
	"github.com/guidoenr/padel-field/controllers"
	_ "github.com/guidoenr/padel-field/controllers"
	"github.com/guidoenr/padel-field/models"
	"github.com/uptrace/bun"
)

func main() {
	//models.InitDB()
	//turnero.InitializeTurnos()
	turno := models.Turno{
		BaseModel: bun.BaseModel{},
		ID:        2,
		OwnerId:   1,
	}
	controllers.ReserveTurno(&turno, 0)
}
