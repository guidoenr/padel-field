package controllers

import (
	"context"
	"fmt"
	"github.com/guidoenr/padel-field/models"
)

// GetAvailableTurnos returns the entire list of available turnos of each weekday
// TODO -> think if its better to filter per/day in the SQL query, in the front or in golang?
func GetAvailableTurnos() ([]models.Turno, error) {
	var availableTurnos []models.Turno

	// initialize the DB cursor
	db := models.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	_, err := db.NewSelect().
		Model(&availableTurnos).
		Where("status = ?", "DISPONIBLE").
		//Where("day = ?", "LUNES").
		ScanAndCount(context.Background())
	if err != nil {
		fmt.Printf("error getting available turnos: %v", err)
	}

	fmt.Println(availableTurnos)

	defer db.Close()
	return availableTurnos, err
}

// CancelTurno changes the state of the turno to AVAILABLE and set the OwnerId = 0
// which means that turno is not related to any user
func CancelTurno(turno *models.Turno) error {
	turno.OwnerId = 0
	turno.Status = models.AVAILABLE

	db := models.InitDB()
	_, err := db.NewUpdate().
		Model(turno).
		Column("status", "owner").
		WherePK().
		Exec(context.Background())

	if err != nil {
		fmt.Printf("error persisting turno: %v", err)
	}
	fmt.Printf("modified turno: %v\n", turno.String())

	defer db.Close()
	return err
}

// ReserveTurno changes the status of the turno to RESERVED and set the ownerId of the user
// that made the request of reserve
func ReserveTurno(turno *models.Turno, ownerId int64) error {
	turno.OwnerId = ownerId
	turno.Status = models.RESERVERD

	db := models.InitDB()
	_, err := db.NewUpdate().
		Model(turno).
		Column("status", "owner").
		WherePK().
		Exec(context.Background())

	if err != nil {
		fmt.Printf("error persisting turno: %v", err)
	}

	defer db.Close()
	return err
}

func persistTurno(turno *models.Turno) error {
	db := models.InitDB()
	_, err := db.NewInsert().
		Model(turno).
		Exec(context.Background())

	if err != nil {
		fmt.Printf("error persisting turno: %v", err)
	}
	fmt.Printf("added turno: %v\n", turno.String())

	defer db.Close()
	return err
}
