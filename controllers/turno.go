package controllers

import (
	"context"
	"fmt"
	"github.com/guidoenr/padel-field/models"
)

func ReserveTurno(turno *models.Turno, ownerId int64) {
	turno.OwnerId = ownerId
	turno.Status = models.Reserved

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
}

func PersistTurno(turno *models.Turno) {
	db := models.InitDB()
	_, err := db.NewInsert().
		Model(turno).
		Exec(context.Background())

	if err != nil {
		fmt.Printf("error persisting turno: %v", err)
	}
	fmt.Printf("added turno: %v\n", turno.String())

	defer db.Close()
}
