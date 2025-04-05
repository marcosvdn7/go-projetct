package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func initDBConnection() (*sql.DB, error) {
	logger := GetLogger("db")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	logger.Info(sql.Drivers())
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Infof("Successfully connected do database %s", dbname)

	return db, nil
}
