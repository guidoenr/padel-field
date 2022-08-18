package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	// TODO
	uri string = "postgres://guido:guido@127.0.0.1:5432/padelfield?sslmode=disable"
)

func Init() {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatalf("Error loading the db: %v", err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
