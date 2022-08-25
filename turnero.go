package main

import (
	"context"
	"fmt"
	"github.com/guidoenr/padel-field/controllers"
	"github.com/guidoenr/padel-field/models"
	"time"
)

// InitializeTurnos [WARNING] it's a zero-day kind of function, it loads all the available turnos into the db
// and clean the existing ones, put all the turnos on available status
func InitializeTurnos() {
	var turnos []models.Turno

	err := cleanDB()
	if err != nil {
		fmt.Printf("error cleaning the db: %v", err)
	}

	initialHour := 9 // that means the turnos start from 09:00 am
	today := time.Now()

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
		turno.Date, _ = time.Parse("02-01-2006", newDate.Format("02-01-2006"))
		turno.Status = models.AVAILABLE
		turno.Day = getWeekDay(newDate)
		// for each turno in a day (time range from 09:00 to 23:00)
		for j := 0; j < 15; j++ {
			t2 := time.Date(today.Year(), today.Month(), today.Day(), initialHour+j, 0, 0, 0, today.Location())
			turno.Hour = t2.Hour()
			turnos = append(turnos, turno)
		}
	}

	for _, t := range turnos {
		err := controllers.PersistTurno(&t)
		if err != nil {
			fmt.Printf("error persisting turno: %v", err)
		}
	}

}

// UpdateTurnos will update all the turnos that are older than today's date
func UpdateTurnos() {
	var outdatedTurnos []models.Turno
	todayDate, _ := time.Parse("02-01-2006", time.Now().Format("02-01-2006"))
	db := models.InitDB()

	// select all the turnos where the date is older than today
	count, err := db.NewSelect().
		Model(&outdatedTurnos).
		Where("date < ?", todayDate). // be aware of handling the = in >= or <=
		ScanAndCount(context.Background())

	defer db.Close()

	if err != nil {
		fmt.Printf("error getting turnos out of date: %v", err)
	}

	fmt.Printf("turnos out of date: %d \n", count)
	fmt.Println(outdatedTurnos)
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
