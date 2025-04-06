package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvdn7/go-projetct/cmd/api/config"
)

func InitializeRouter() {
	logger := config.GetLogger("router")
	router := gin.Default()
	routerGroup := router.Group("go-project/api")

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	InitializeCharacterRoutes(routerGroup)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		logger.Errorf("Failed to initiate server: %v", err)
		panic(err)
	}
}
