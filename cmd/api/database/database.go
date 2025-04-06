package database

import (
	"database/sql"
	"github.com/marcosvdn7/go-projetct/cmd/api/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func InitializeDatabase() {
	db = config.GetDB()
	logger = config.GetLogger("database")
}
