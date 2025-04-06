package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvdn7/go-projetct/cmd/api/handler"
)

func InitializeCharacterRoutes(routerGroup *gin.RouterGroup) {
	{
		routerGroup.POST("/character", handler.CreateCharacterHandler)
		routerGroup.GET("/character", handler.GetCharacterHandler)
		routerGroup.PUT("/character/", handler.UpdateCharacterHandler)
		routerGroup.DELETE("/character", handler.DeleteCharacterHandler)
		routerGroup.GET("/characters", handler.ListCharactersHandler)
	}
}
