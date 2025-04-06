package handler

import (
	"github.com/marcosvdn7/go-projetct/cmd/api/config"
)

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
}
