package model

type Music struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Artist string `json:"artist" db:"artist"`
	Url    string `json:"url" db:"url"`
	Cover  string `json:"cover" db:"cover"`
}
