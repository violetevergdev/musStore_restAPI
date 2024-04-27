package app

import "github.com/gin-gonic/gin"

func Run() {
	Router:= gin.Default()

	Router.GET("/app",  func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	
	})

	Router.Run("localhost:8080")
}