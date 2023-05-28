package model

import "time"

type ElectionDetailed struct {
	SubjectId      string `json:"subject_id" db:"subject_id"`
	SubjectName    string `json:"subject_name" db:"subject_name"`
	Teacher        string `json:"teacher" db:"teacher"`
	Classification string `json:"classification" db:"classification"`
	Credit         string `json:"credit" db:"credit"`
	College        string `json:"college" db:"college"`
	Attendance     string `json:"attendance" db:"attendance"`
	Score          string `json:"score" db:"score"`
	Evaluation     string `json:"evaluation" db:"evaluation"`
}

type Election struct {
	SubjectId   string `json:"subject_id" db:"subject_id"`
	SubjectName string `json:"subject_name" db:"subject_name"`
	Teacher     string `json:"teacher" db:"teacher"`
}

type ElectionComment struct {
	Id          int       `json:"id" db:"id"`
	SubjectId   string    `json:"subject_id" db:"subject_id"`
	SubjectName string    `json:"subject_name" db:"subject_name"`
	CreateTime  time.Time `db:"create_time" json:"createTime"`
	Comment     string    `json:"comment" db:"comment"`
}
