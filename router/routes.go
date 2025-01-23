package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	router.Use(func(context *gin.Context) {
		if context.Query("status") != "open" {
			context.JSON(400, gin.H{
				"message": "API is closed at the moment. Please check back later. Thank you!",
			})
			context.Abort()
			return
		}
		context.Next()
	})
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Opening API",
			})
		})
		v1.POST("/posting", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Posting API",
			})
		})
		v1.DELETE("/deleting", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Deleting API",
			})
		})
		v1.PUT("/puting", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Puting API",
			})
		})
		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "openings API",
			})
		})
	}
}
