package models

type Track struct {
	ID     uint   `json:"id"`
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}