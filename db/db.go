package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/guidoenr/padel-field/models"
	_ "github.com/lib/pq"
)

const (
	// TODO
	uri string = "postgres://guido:guido@127.0.0.1:5432/padelfield?sslmode=disable"
)

func Init() {
	db := pg.Connect(&pg.Options{
		Addr: uri,
	})

	err := db.Ping(context.Background())
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer db.Close()
	err = createSchema(db)

}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Turno)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
