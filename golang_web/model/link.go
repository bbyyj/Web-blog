package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Link struct {
	Id         int    `json:"id" db:"id" gorm:"primary_key"`
	Name       string `json:"name" db:"name"`
	Desc       string `json:"desc" db:"desc"`
	Url        string `json:"url" db:"url"`
	CategoryId int    `json:"categoryId" db:"category_id"`
	Icon       string `json:"icon" db:"icon"`
}

type LinkCategory struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/*
type ResourceManage struct {
	Id          int        `json:"id" db:"id" gorm:"primary_key"`
	Name        string     `json:"name" db:"name"`
	Desc        string     `json:"desc" db:"desc"`
	Url         string     `json:"url" db:"url"`
	CategoryId  int        `json:"categoryid" db:"category_id"`
	DownloadNum int        `json:"downloadnum" db:"downloadnum"`
	FileSize    int        `json:"filesize" db:"filesize"`
	CreatedAt   time.Time  `json:"createdat" db:"createdat"`
	UpdatedAt   time.Time  `json:"updatedat" db:"updatedat"`
	DeletedAt   *time.Time `sql:"index" json:"deletedat" db:"deletedat"`
}*/

type ResourceManage struct {
	gorm.Model
	Name        string `json:"name" db:"name"`
	Desc        string `json:"desc" db:"desc"`
	Url         string `json:"url" db:"url"`
	CategoryId  int    `json:"categoryid" db:"category_id"`
	DownloadNum int    `json:"downloadnum" db:"downloadnum"`
	FileSize    int    `json:"filesize" db:"filesize"`
}

type ResourceUrl struct {
	Name string `json:"name" db:"name"`
	Url  string `json:"url" db:"url"`
}

type ResourceList struct {
	Name        string    `json:"name" db:"name"`
	Desc        string    `json:"desc" db:"desc"`
	CategoryId  int       `json:"categoryid" db:"category_id"`
	DownloadNum int       `json:"downloadnum" db:"downloadnum"`
	FileSize    int       `json:"filesize" db:"filesize"`
	CreatedAt   time.Time `json:"createdat" db:"createdat"`
}

type ResResult struct {
	Data    []interface{} `json:"data"`
	Status  uint32        `json:"status"`
	Message string        `json:"message"`
}

/*
	type ResoureDetail struct {
		Id          int        `json:"id" db:"id" gorm:"primary_key"`
		Name        string     `json:"name" db:"name"`
		Desc        string     `json:"desc" db:"desc"`
		Url         string     `json:"url" db:"url"`
		CategoryId  int        `json:"categoryId" db:"category_id"`
		DownloadNum int        `json:"downloadnum" db:"downloadnum"`
		FileSize    int        `json:"filesize" db:"filesize"`
		CreatedAt   time.Time  `json:"createdat" db:"createdat"`
		UpdatedAt   time.Time  `json:"updatedat" db:"updatedat"`
		DeletedAt   *time.Time `sql:"index" json:"deletedat" db:"deletedat"`
	}
*/
//var FileRoot = "D:/Go project/Go_UPandDownload/downloads/%s"
var FileRoot = "./downloads/%s"
