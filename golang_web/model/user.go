package model

type User struct {
	Id       int    `db:"id" json:"id"`
	Nickname string `db:"nickname" json:"nickname"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Avatar   string `db:"avatar" json:"avatar"`
}
