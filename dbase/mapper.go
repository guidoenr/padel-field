package dbase

import (
	"context"
	"fmt"
	"github.com/guidoenr/padel-field/models"
	"github.com/uptrace/bun"
)

func PersistTurno(db *bun.DB, turno *models.Turno) {
	result, err := db.NewInsert().Model(turno).Exec(context.Background())
	if err != nil {
		fmt.Printf("error persisting turno: %v", err)
	}
	fmt.Println(result)
}

//
//func updateTurno(dbase *bun.DB, turno models.Turno, newStatus string, newOwnerId int) {
//
//	// update the turno status
//	res, _ := dbase.NewUpdate().
//		Model(turno).
//		WherePK().
//		Set("status = _status").
//		Exec(context.Background())
//}
