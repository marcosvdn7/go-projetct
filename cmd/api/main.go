package main

import (
	"github.com/marcosvdn7/go-projetct/cmd/api/config"
	"github.com/marcosvdn7/go-projetct/cmd/api/database"
	"github.com/marcosvdn7/go-projetct/cmd/api/handler"
	"github.com/marcosvdn7/go-projetct/cmd/api/router"
)

func main() {
	config.Init()
	handler.InitializeHandler()
	database.InitializeDatabase()
	router.InitializeRouter()
}
