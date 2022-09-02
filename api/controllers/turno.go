package controllers

import (
	"context"
	"github.com/guidoenr/padel-field/api/errs"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"github.com/guidoenr/padel-field/models/psdb"
	"github.com/guidoenr/padel-field/tools"
	"strconv"
	"strings"
)

// GetAvailableTurnos returns the entire list of available turnos of each weekday
func GetAvailableTurnos() ([]models.Turno, errs.RequestError) {
	var availableTurnos []models.Turno
	logger.Loginfo.Println("getting availabale turnos")
	// initialize the DB cursor
	db := psdb.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	err := db.NewSelect().
		Model(&availableTurnos).
		Where("status = ?", "DISPONIBLE").
		Scan(context.Background())

	if err != nil {
		return []models.Turno{}, errs.ThrowError(err, "can not get available turnos", -1)
	}

	defer db.Close()
	return availableTurnos, errs.ThrowEmptyError()
}

// GetAvailableTurnosByDay returns the entire list of available turnos by a day
func GetAvailableTurnosByDay(day string) ([]models.Turno, errs.RequestError) {
	var availableTurnos []models.Turno
	logger.Loginfo.Printf("getting availabale turnos for day: %s", day)
	// initialize the DB cursor
	db := psdb.InitDB()

	today, _ := tools.GetTodayDate()
	oneWeek := today.AddDate(0, 0, 7)
	// select * from turnos where status = "DISPONIBLE"
	err := db.NewSelect().
		Model(&availableTurnos).
		Where("status = ?", "DISPONIBLE").
		Where("day = ?", strings.ToUpper(day)).
		Where("date < ?", oneWeek).
		Scan(context.Background())

	if err != nil {
		return []models.Turno{}, errs.ThrowError(err, "can not get available turnos for this day", -1)
	}

	defer db.Close()
	return availableTurnos, errs.ThrowEmptyError()
}

// GetTurnosByOwnerId returns all the turnos that are linked to a specific user
func GetTurnosByOwnerId(ownerId string) ([]models.Turno, errs.RequestError) {
	var turnosByOwner []models.Turno
	logger.Loginfo.Printf("getting turnos for user '%d'", ownerId)
	// initialize the DB cursor
	db := psdb.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	err := db.NewSelect().
		Model(&turnosByOwner).
		Where("owner = ?", ownerId).
		Scan(context.Background())

	if err != nil {
		return []models.Turno{}, errs.ThrowError(err, "can not get turnos for that user", -1)
	}

	defer db.Close()
	return turnosByOwner, errs.ThrowEmptyError()
}

// GetTurnoById returns one single turno linked to the turnoId
func GetTurnoById(id string) (models.Turno, errs.RequestError) {
	var turnoById models.Turno
	idConv, _ := strconv.Atoi(id)

	logger.Loginfo.Printf("finding turno by id: '%d'", id)
	// initialize the DB cursor
	db := psdb.InitDB()

	err := db.NewSelect().
		Model(&turnoById).
		Where("id = ?", idConv).
		Scan(context.Background())

	if err != nil {
		return models.Turno{}, errs.ThrowError(err, "turno not found", -1)
	}

	defer db.Close()
	return turnoById, errs.ThrowEmptyError()
}

// CancelTurno changes the state of the turno to AVAILABLE and set the OwnerId = 0
// which means that turno is not related to any user
func CancelTurno(id string, ownerId int64) errs.RequestError {
	logger.Loginfo.Println("canceling turno")
	db := psdb.InitDB()

	turno := new(models.Turno)

	_, err := db.NewUpdate().
		Model(turno).
		Set("status = ?", models.AVAILABLE).
		Where("id = ?", id).
		Where("owner = ?", ownerId).
		Exec(context.Background())

	if err != nil {
		return errs.ThrowError(err, "can not cancel the turno", -1)
	}

	defer db.Close()
	return errs.ThrowEmptyError()
}

// ReserveTurno changes the status of the turno to RESERVED and set the ownerId of the user
// that made the request of reserve
func ReserveTurno(id string, ownerId int64) errs.RequestError {
	logger.Loginfo.Println("reserving turno")
	db := psdb.InitDB()

	turno := new(models.Turno)

	_, err := db.NewUpdate().
		Model(turno).
		Set("status = ?", models.RESERVERD).
		Set("owner = ?", ownerId).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		return errs.ThrowError(err, "can not get reserve turno", -1)
	}

	defer db.Close()
	return errs.ThrowEmptyError()
}

// PersistTurnos is a built-in function to map turnos into the db
func PersistTurnos(turnos []models.Turno) errs.RequestError {
	db := psdb.InitDB()
	_, err := db.NewInsert().
		Model(&turnos).
		Exec(context.Background())

	if err != nil {
		return errs.ThrowError(err, "can not persist turnos", -1)
	}

	defer db.Close()
	return errs.ThrowEmptyError()
}
