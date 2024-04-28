package app

import (
	"musicstore_rest_api/internal/database/postgres"
	"musicstore_rest_api/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	postgres.ConnectDB()
}

func Run() {
	Router:= gin.Default()

	Router.GET("/app", models.GetTracks )
	Router.POST("/app", models.CreateTrack)

	Router.Run("localhost:8080")
}