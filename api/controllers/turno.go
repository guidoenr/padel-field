package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"github.com/guidoenr/padel-field/tools"
	"strconv"
	"strings"
)

// GetAvailableTurnos returns the entire list of available turnos of each weekday
func GetAvailableTurnos() ([]models.Turno, error) {
	var availableTurnos []models.Turno
	logger.Loginfo.Println("getting availabale turnos")
	// initialize the DB cursor
	db := models.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	err := db.NewSelect().
		Model(&availableTurnos).
		Where("status = ?", "DISPONIBLE").
		Scan(context.Background())
	if err != nil {
		logger.Logerror.Printf("error getting available turnos: %v", err)
	}

	defer db.Close()
	return availableTurnos, err
}

// GetAvailableTurnosByDay returns the entire list of available turnos by a day
func GetAvailableTurnosByDay(day string) ([]models.Turno, error) {
	var availableTurnos []models.Turno
	logger.Loginfo.Printf("getting availabale turnos for day: %s", day)
	// initialize the DB cursor
	db := models.InitDB()

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
		logger.Logerror.Printf("error getting available turnos: %v", err)
	}

	defer db.Close()
	return availableTurnos, err
}

// GetTurnosByOwnerId returns all the turnos that are linked to a specific user
func GetTurnosByOwnerId(ownerId string) ([]models.Turno, error) {
	var turnosByOwner []models.Turno
	logger.Loginfo.Printf("getting turnos for user '%d'", ownerId)
	// initialize the DB cursor
	db := models.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	err := db.NewSelect().
		Model(&turnosByOwner).
		Where("owner = ?", ownerId).
		Scan(context.Background())

	if err != nil {
		logger.Logerror.Printf("getting turnos for user '%d'", ownerId)
	}

	defer db.Close()
	return turnosByOwner, err
}

// GetTurnoById returns one single turno linked to the turnoId
func GetTurnoById(id string) (models.Turno, error) {
	var turnoById models.Turno
	idConv, _ := strconv.Atoi(id)

	logger.Loginfo.Printf("finding turno by id: '%d'", id)
	// initialize the DB cursor
	db := models.InitDB()

	err := db.NewSelect().
		Model(&turnoById).
		Where("id = ?", idConv).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("turno '%s' not found", id)
		logger.Logerror.Println(msg)
		err = errors.New(msg)
	}

	defer db.Close()
	return turnoById, err
}

// CancelTurno changes the state of the turno to AVAILABLE and set the OwnerId = 0
// which means that turno is not related to any user
func CancelTurno(id string, ownerId int64) error {
	logger.Loginfo.Println("canceling turno")
	db := models.InitDB()

	turno := new(models.Turno)

	_, err := db.NewUpdate().
		Model(turno).
		Set("status = ?", models.AVAILABLE).
		Where("id = ?", id).
		Where("owner = ?", ownerId).
		Exec(context.Background())

	if err != nil {
		msg := fmt.Sprintf("canceling turno '%s' for user '%d': %v \n", id, ownerId, err)
		logger.Logerror.Println(msg)
		err = errors.New(msg)
	}

	defer db.Close()
	return err
}

// ReserveTurno changes the status of the turno to RESERVED and set the ownerId of the user
// that made the request of reserve
func ReserveTurno(id string, ownerId int64) error {
	logger.Loginfo.Println("reserving turno")
	db := models.InitDB()

	turno := new(models.Turno)

	_, err := db.NewUpdate().
		Model(turno).
		Set("status = ?", models.RESERVERD).
		Set("owner = ?", ownerId).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		msg := fmt.Sprintf("reserving turno '%s' for user '%d': %v \n", id, ownerId, err)
		logger.Logerror.Println(msg)
		err = errors.New(msg)
	}

	defer db.Close()
	return err
}

// PersistTurnos is a built-in function to map turnos into the db
func PersistTurnos(turnos []models.Turno) error {
	db := models.InitDB()
	_, err := db.NewInsert().
		Model(&turnos).
		Exec(context.Background())

	if err != nil {
		logger.Logerror.Printf("error persisting turnos: %v", err)
	}

	defer db.Close()
	return err
}