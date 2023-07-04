package model

import "time"

type Motto struct {
	Id         uint32    `db:"id" json:"id"`
	Ch         string    `db:"ch" json:"ch"`
	En         string    `db:"en" json:"en"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}
