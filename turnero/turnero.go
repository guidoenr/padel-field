package turnero

import (
	"github.com/guidoenr/padel-field/controllers"
	"github.com/guidoenr/padel-field/models"
	"time"
)

// InitializeTurnos it's a zero-day kind of function, it loads all the available turnos into the db
func InitializeTurnos() {
	var turnos []models.Turno
	initialHour := 9
	today := time.Now()
	for i := 0; i <= 7; i++ {
		var turno models.Turno
		turno.Field = "blindex"
		turno.Date = today.AddDate(0, 0, i)
		turno.Status = models.AVAILABLE
		turno.Day = getWeekDay(turno.Date)
		for j := 0; j <= 14; j++ {
			t2 := time.Date(today.Year(), today.Month(), today.Day(), initialHour+j, 0, 0, 0, today.Location())
			turno.Hour = t2.Hour()
			turnos = append(turnos, turno)
		}
	}

	for _, t := range turnos {
		controllers.PersistTurno(&t)
	}

}

// daysMap a map with the purpose of be a dictionary of translations
var daysMap = map[string]string{
	"Monday":    "LUNES",
	"Tuesday":   "MARTES",
	"Wednesday": "MIERCOLES",
	"Thursday":  "JUEVES",
	"Friday":    "VIERNES",
	"Saturday":  "SABADO",
	"Sunday":    "DOMINGO",
}

// getWeekDay is a basic transform-language date function that transform the datetime weekday
// english string into a spanish string (e.g Monday -> LUNES)
func getWeekDay(datetime time.Time) string {
	return daysMap[datetime.Weekday().String()]
}
