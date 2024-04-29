package app

import (
	"musicstore_rest_api/internal/database/postgres"
	"musicstore_rest_api/internal/transport/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	postgres.ConnectDB()
}

func Run() {
	defer postgres.ConnectDB().Close()

	Router:= gin.Default()

	Router.GET("/app", rest.GetTracks )
	Router.GET("/app/:id", rest.GetTrackById)
	Router.POST("/app", rest.CreateTrack)

	Router.Run("localhost:8080")
}