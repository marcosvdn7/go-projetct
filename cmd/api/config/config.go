package config

import (
	"database/sql"
	"errors"
	"github.com/marcosvdn7/go-projetct/cmd/api/router"
)

var (
	logger *Logger
	db     *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

func GetLogger(prefix string) *Logger {
	logger = newLogger(prefix)
	return logger
}

func InitializeDB() (db *sql.DB, err error) {
	db, err = initDBConnection()
	if err != nil {
		return nil, errors.New("Error connecting to database: " + err.Error())
	}

	return db, nil
}

func InitializeServer() {
	if err := router.InitializeRouter(); err != nil {
		logger.Errorf("Error initializing routes %v", err)
		panic(err)
	}
}
