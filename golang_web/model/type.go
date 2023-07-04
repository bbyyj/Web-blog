package model

type TheType struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Count int    `db:"count" json:"count"`
}
