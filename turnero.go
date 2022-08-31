package main

import (
	"context"
	"github.com/guidoenr/padel-field/api/controllers"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"github.com/guidoenr/padel-field/tools"
	"time"
)

// -------------------------- TURNOS

// InitializeTurnos it's a zero-day kind of function, it loads all the available turnos into the db
// and clean the existing ones, put all the turnos on available status
func InitializeTurnos() error {
	logger.Logwarning.Println("initialazing turnos..")
	var turnos []models.Turno

	// turnos range from 9 -> 23 and in one month (30)
	// [TODO, maybe create a config file]
	initialHour := 9
	lastHour := 23
	daysOfTurno := 30

	today, location := tools.GetTodayDate()

	// for each weekday
	for i := 0; i < daysOfTurno; i++ {
		var turno models.Turno
		turno.Field = "blindex"
		// every time that we iterate, we plus one day to the day
		// (e.g) if today is 24 of July, in the second iteration it will be 25 of July and then
		newDate := today.AddDate(0, 0, i)
		turno.Date, _ = time.ParseInLocation("02-01-2006", newDate.Format("02-01-2006"), location)
		turno.Status = models.AVAILABLE
		turno.Day = getWeekDay(newDate)
		// for each turno in a day (time range from 09:00 to 23:00)
		for j := initialHour; j <= lastHour; j++ {
			turno.Hour = j
			turnos = append(turnos, turno)
		}
	}

	err := controllers.PersistTurnos(turnos)
	if err != nil {
		logger.Logerror.Printf("initializing turnos: %v", err)
	}
	return err
}

// UpdateTurnos will update all the turnos that are older than today's date
func UpdateTurnos() {
	//var outdatedTurnos []models.Turno
	logger.Loginfo.Println("Updating out-of-date turnos")
	todayDate, _ := tools.GetTodayDate()
	db := models.InitDB()

	turno := new(models.Turno)

	// updating the turnos that are not fijos, setting the status to available
	_, err := db.NewUpdate().
		Model(turno).
		Set("date = ?", todayDate.AddDate(0, 0, 30)).
		Set("status = ?", models.AVAILABLE).
		Where("fijo = ?", false).
		Where("date < ?", todayDate).
		Exec(context.Background())

	// updating the date of the turnos that are fijos
	_, err = db.NewUpdate().
		Model(turno).
		Set("date = ?", todayDate.AddDate(0, 0, 30)).
		Where("fijo = ?", true).
		Where("date < ?", todayDate).
		Exec(context.Background())

	defer db.Close()

	if err != nil {
		logger.Logerror.Printf("error updating outdated turnos: %v", err)
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

// -------------------------- USERS

func InitializeUsers() error {
	logger.Logwarning.Println("initialazing users..")
	db := models.InitDB()
	rootUser := models.User{
		Username:  "root",
		Password:  "rootoor",
		Role:      models.ADMIN,
		Phone:     "",
		Email:     "",
		Firstname: "Root",
		Lastname:  "Root",
	}

	_, err := db.NewInsert().
		Model(&rootUser).
		Exec(context.Background())

	defer db.Close()
	return err
}
