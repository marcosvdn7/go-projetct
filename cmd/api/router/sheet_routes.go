package router

import "github.com/gin-gonic/gin"

func initializeSheetRoutes(routerGroup *gin.RouterGroup) {
	{
		routerGroup.POST("/sheet", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Created",
			})
		})
		routerGroup.GET("/sheet")
		routerGroup.PUT("/sheet")
		routerGroup.DELETE("/sheet")
		routerGroup.GET("/sheets")
	}
}
