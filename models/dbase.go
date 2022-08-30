package models

import (
	"context"
	"database/sql"
	"github.com/guidoenr/padel-field/logger"
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
		logger.Logerror.Printf("error making Ping() to dbase: %v", err)
	}
	logger.Loginfo.Printf("connected to the db: %v \n", db.DB.Stats())

	return db
}

// RestartDb drop the tables and creates database schema for User and Turno models.
func RestartDb(db *bun.DB) error {
	var err error
	err = nil
	logger.Logwarning.Println("RESTARTING DB..")
	_, err = db.Query("DROP TABLE turnos;")
	_, err = db.Query("DROP TABLE users;")
	err = createSchemas(db)
	return err
}

// CreateSchemas create all the tables of the database following the bun models
func createSchemas(db *bun.DB) error {
	var err error
	_, err = db.NewCreateTable().
		Model((*User)(nil)).
		Exec(context.Background())

	_, err = db.NewCreateTable().
		Model((*Turno)(nil)).
		Exec(context.Background())

	return err
}

// loadDBConnector load the environment variables from `local.env` // TODO, change later to heroku maybe?
func loadDBConnector() *pgdriver.Connector {
	err := godotenv.Load("models/local.env")
	if err != nil {
		logger.Logerror.Printf("error reading 'local.env' file err: %v \n", err)
	}

	user := os.Getenv("DB_USER")
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
