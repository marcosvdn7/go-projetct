package main

import (
	"database/sql"
	"github.com/marcosvdn7/go-projetct/cmd/api/config"
	"os/exec"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func main() {
	logger = config.GetLogger("main")

	cmd := exec.Command("docker", "compose", "up", "-d", "go-project")
	cmd.Dir = "C:\\git\\go-projetct\\"
	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error starting docker db: %v", err)
		panic(err)
	}

	db, err = config.InitializeDB()
	if err != nil {
		logger.Error(err)
		return
	}

	config.InitializeServer()
}
