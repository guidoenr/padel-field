package main

import (
	_ "github.com/guidoenr/padel-field/controllers"
	"github.com/guidoenr/padel-field/models"
)

func main() {
	//models.InitDB()
	//turnero.InitializeTurnos()
	//turno := models.Turno{
	//	BaseModel: bun.BaseModel{},
	//	ID:        2,
	//	OwnerId:   1,
	//}
	//controllers.ReserveTurno(&turno, 0)
	models.InitDB()
}
