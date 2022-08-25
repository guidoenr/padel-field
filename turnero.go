package main

import (
	"context"
	"github.com/guidoenr/padel-field/controllers"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"time"
)

// InitializeTurnos it's a zero-day kind of function, it loads all the available turnos into the db
// and clean the existing ones, put all the turnos on available status
func InitializeTurnos() {
	logger.Logwarning.Println("Initialazing turnos..")
	var turnos []models.Turno

	err := cleanDB()
	if err != nil {
		logger.Logerror.Printf("error cleaning the db: %v", err)
	}

	initialHour := 9 // that means the turnos start from 09:00 am
	lastHour := 23
	today := getTodayDate()
	// daysOfTurno are all the turnos that we are going to load
	// inside the database
	daysOfTurno := 30

	// for each weekday
	for i := 0; i < daysOfTurno; i++ {
		var turno models.Turno
		turno.Field = "blindex"
		// every time that we iterate, we plus one day to the day
		// (e.g) if today is 24 of July, in the second iteration it will be 25 of July and then
		newDate := today.AddDate(0, 0, i)
		turno.Date, _ = time.ParseInLocation("02-01-2006", newDate.Format("02-01-2006"), time.Local)
		turno.Status = models.AVAILABLE
		turno.Day = getWeekDay(newDate)
		// for each turno in a day (time range from 09:00 to 23:00)
		for j := initialHour; j <= lastHour; j++ {
			turno.Hour = j
			turnos = append(turnos, turno)
		}
	}

	for _, t := range turnos {
		err = controllers.PersistTurno(&t)
		if err != nil {
			logger.Logerror.Printf("error persisting turno: %v", err)
		}
	}

}

// UpdateTurnos will update all the turnos that are older than today's date
func UpdateTurnos() {
	//var outdatedTurnos []models.Turno
	logger.Loginfo.Println("Updating out-of-date turnos")
	todayDate := getTodayDate()
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

// cleanDb drop al the tables of the db
func cleanDB() error {
	db := models.InitDB()
	err := models.RestartDb(db)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

// getTodayDate returns the current local datetime in America/Buenos_Aires
func getTodayDate() time.Time {
	location, err := time.LoadLocation("America/Buenos_Aires")
	if err != nil {
		logger.Logerror.Printf("error loading the location: %v", err)
	}
	todayDate, err := time.ParseInLocation("02-01-2006", time.Now().Format("02-01-2006"), location)
	if err != nil {
		logger.Logerror.Printf("error getting localTime: %v", err)
	}
	return todayDate
}
