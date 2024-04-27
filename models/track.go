package models

type Track struct {
	ID     uint   `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}