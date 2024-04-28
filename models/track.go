package models

import (
	"net/http"

	"musicstore_rest_api/internal/database/postgres"

	"github.com/gin-gonic/gin"
)

type Track struct {
	ID     uint   `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetTracks(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var tracks []Track

	rows, err := postgres.ConnectDB().Query("select * from traks")
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Internal server error",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t Track

		err = rows.Scan(&t.ID, &t.Artist, &t.Title)
		if err!= nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"message": "Internal server error",
			})
			return
		}

		tracks = append(tracks, t)
	}
	ctx.JSON(http.StatusOK, &tracks)
}

func CreateTrack(ctx *gin.Context) {
	var track Track

	err := ctx.Bind(&track); if err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "Bad request",
		})
		return
	}

	stmt, err := postgres.ConnectDB().Prepare("insert into traks (artist, title) values ($1, $2)"); if err!= nil {ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Internal server error",
		})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(track.Artist, track.Title); if err!= nil {ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
			"message": "Internal server error",
		})
	}
	

	ctx.JSON(http.StatusCreated, track)
}