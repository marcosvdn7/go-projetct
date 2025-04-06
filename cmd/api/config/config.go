package config

import (
	"database/sql"
	"os/exec"
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

func Init() {
	logger = newLogger("config")
	initializeDB()
}

func GetLogger(p string) *Logger {
	return newLogger(p)
}

func initializeDB() {
	initializeDocker()

	var err error
	db, err = initDBConnection()
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func initializeDocker() {
	cmd := exec.Command("docker", "compose", "up", "-d", "go-project")
	cmd.Dir = "C:\\git\\go-projetct\\"
	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error starting docker db: %v", err.Error())
		panic(err)
	}
}
