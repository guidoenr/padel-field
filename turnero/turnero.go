package turnero

import (
	"fmt"
	"github.com/guidoenr/padel-field/models"
	"time"
)

var daysMap = map[string]string{
	"Monday":    "LUNES",
	"Tuesday":   "MARTES",
	"Wednesday": "MIERCOLES",
	"Thursday":  "JUEVES",
	"Friday":    "VIERNES",
	"Saturday":  "SABADO",
	"Sunday":    "DOMINGO",
}

func InitializeTurnos() {

	var turnos []models.Turno
	initialHour := 9
	today := time.Now()
	for i := 0; i <= 7; i++ {
		var turno models.Turno
		turno.Field = "blindex"
		turno.Date = today.AddDate(0, 0, i)
		turno.Status = models.Available
		turno.Day = getWeekDay(turno.Date)
		for j := 0; j <= 15; j++ {
			t2 := time.Date(today.Year(), today.Month(), today.Day(), initialHour+j, 0, 0, 0, today.Location())
			turno.Hour = t2.Hour()
			turnos = append(turnos, turno)
		}
	}

	for _, t := range turnos {
		fmt.Println(t.String())
	}
}

func getWeekDay(datetime time.Time) string {
	return daysMap[datetime.Weekday().String()]
}
