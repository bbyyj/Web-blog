package model

type Tag struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	count int    `db:"count" json:"count"`
}

// 多对多 关联表
type BlogTag struct {
	BlogId int `db:"blog_id"`
	TagId  int `db:"tag_id"`
}
