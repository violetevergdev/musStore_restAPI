package rest

import (
	"musicstore_rest_api/internal/database/postgres"
	"musicstore_rest_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: БЫли вынесены контроллеры
func GetTracks(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var tracks []models.Track

	rows, err := postgres.ConnectDB().Query("select * from traks")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Internal server error",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Track

		err = rows.Scan(&t.ID, &t.Artist, &t.Title)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Internal server error",
			})
			return
		}

		tracks = append(tracks, t)
	}
	ctx.JSON(http.StatusOK, &tracks)
}

func GetTrackById(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var track models.Track

	id := ctx.Param("id")
	err := postgres.ConnectDB().QueryRow("select * from traks where id=$1", id).Scan(&track.ID, &track.Artist, &track.Title) 
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "This row does not exist",
		})
		return
	}

	ctx.JSON(http.StatusOK, &track)

}

func CreateTrack(ctx *gin.Context) {
	var track models.Track

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

// func UpdateTrack(ctx *gin.Context) {
// 	 var track models.Track
// 	 id := ctx.Param("id")

// 	 err := ctx.Bind(&track); if err!= nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			 "error": err.Error(),
// 			 "message": "Bad request",
// 		 })
// 		 return
// 	 }

// 	 stmt, err := postgres.ConnectDB().Prepare("update traks set artist=$1, title=$2 where id=$3"); if err!= nil {ctx.JSON(http.StatusInternalServerError, gin.H{
// 		"error": err.Error(),
// 			"message": "Internal server error",
// 	 })}
// 	 defer stmt.Close()
// 	 _, err = stmt.Exec(track.Artist, track.Title, id); if err!= nil {ctx.JSON(http.StatusInternalServerError, gin.H{
// 		"error": err.Error(),
// 			"message": "Internal server error",
// 	 })}

// 	 ctx.JSON(http.StatusOK, track)
// }