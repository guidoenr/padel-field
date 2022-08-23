package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

// InitDB only connects to the existing DB | sudo -u postgres psql
func InitDB() *bun.DB {
	// load to connect driver string from config
	pgconn := loadDBConnector()
	sqldb := sql.OpenDB(pgconn)

	// returning the dbase to operate
	db := bun.NewDB(sqldb, pgdialect.New())

	// making ping
	err := db.Ping()
	if err != nil {
		fmt.Printf("error making Ping() to dbase: %v", err)
	}
	fmt.Println("connected succesfully")

	createSchema(db)
	return db
}

// createSchema creates database schema for User and Turno models.
func createSchema(db *bun.DB) {
	// TODO check if the schema exists
	err1, _ := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	err2, _ := db.NewCreateTable().Model((*Turno)(nil)).Exec(context.Background())
	fmt.Println(err1, err2)
}

// loadEnv load the environment variables from `local.env` // TODO, change later to heroku maybe?
func loadDBConnector() *pgdriver.Connector {
	err := godotenv.Load("models/local.env")
	if err != nil {
		fmt.Printf("error reading 'local.env' file err: %v", err)
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")
	addr := os.Getenv("ADDR")

	pgconn := pgdriver.NewConnector(
		pgdriver.WithUser(user),
		pgdriver.WithPassword(password),
		pgdriver.WithDatabase(dbname),
		pgdriver.WithAddr(addr),
		pgdriver.WithNetwork("tcp"),
	)

	return pgconn
}