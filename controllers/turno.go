package controllers

import (
	"context"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
)

// GetAvailableTurnos returns the entire list of available turnos of each weekday
// TODO -> think if is better to filter per/day in the SQL query, in the front or in golang?
// i think that it will be better in SQL, make some test then
func GetAvailableTurnos() ([]models.Turno, error) {
	var availableTurnos []models.Turno
	logger.Loginfo.Println("getting availabale turnos")
	// initialize the DB cursor
	db := models.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	_, err := db.NewSelect().
		Model(&availableTurnos).
		Where("status = ?", "DISPONIBLE").
		//Where("day = ?", "LUNES").
		ScanAndCount(context.Background())
	if err != nil {
		logger.Logerror.Printf("error getting available turnos: %v", err)
	}

	defer db.Close()
	return availableTurnos, err
}

// GetTurnosByOwnerId returns all the turnos that are linked to a specific user
func GetTurnosByOwnerId(ownerId int64) ([]models.Turno, error) {
	var turnosByOwner []models.Turno
	logger.Loginfo.Printf("getting turnos for user '%d'", ownerId)
	// initialize the DB cursor
	db := models.InitDB()

	// select * from turnos where status = "DISPONIBLE"
	_, err := db.NewSelect().
		Model(&turnosByOwner).
		Where("ownerId = ?", ownerId).
		ScanAndCount(context.Background())

	if err != nil {
		logger.Logerror.Printf("error getting turnos for iser '%d'", ownerId)
	}

	defer db.Close()
	return turnosByOwner, err
}

// CancelTurno changes the state of the turno to AVAILABLE and set the OwnerId = 0
// which means that turno is not related to any user
func CancelTurno(turno *models.Turno) error {
	logger.Loginfo.Println("canceling turno")
	turno.OwnerId = 0
	turno.Status = models.AVAILABLE

	db := models.InitDB()
	_, err := db.NewUpdate().
		Model(turno).
		Column("status", "owner").
		WherePK().
		Exec(context.Background())

	if err != nil {
		logger.Logerror.Printf("error canceling turno: %v", err)
	}
	logger.Loginfo.Printf("modified turno: %v\n", turno.String())

	defer db.Close()
	return err
}

// ReserveTurno changes the status of the turno to RESERVED and set the ownerId of the user
// that made the request of reserve
func ReserveTurno(turno *models.Turno, ownerId int64) error {
	turno.OwnerId = ownerId
	turno.Status = models.RESERVERD

	logger.Loginfo.Println("reserving turno")
	db := models.InitDB()
	_, err := db.NewUpdate().
		Model(turno).
		Column("status", "owner").
		WherePK().
		Exec(context.Background())

	if err != nil {
		logger.Logerror.Printf("error reserving turno: %v", err)
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
	logger.Loginfo.Printf("%d turnos persisted", len(turnos))

	defer db.Close()
	return err
}
