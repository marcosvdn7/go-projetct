package router

import "github.com/gin-gonic/gin"

func InitializeRouter() error {
	router := gin.Default()
	routerGroup := router.Group("go-project/api")

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	initializeSheetRoutes(routerGroup)

	return router.Run(":8080")
}
