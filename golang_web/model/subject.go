package model

type Subject struct {
	Id      int    `json:"id" db:"id"`
	Subject string `json:"subject" db:"subject"`
}

type SubjectVue struct {
	Subject []string `json:"subject"`
}
type Chapter struct {
	Id      int    `json:"id" db:"id"`
	Subject string `json:"subject" db:"subject"`
	Chapter string `json:"chapter" db:"chapter"`
}

type ChapterVue struct {
	Subject string   `json:"subject" `
	Chapter []string `json:"chapter" `
}
