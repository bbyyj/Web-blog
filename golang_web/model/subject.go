package model

type Subject struct {
	Id      int    `json:"id" db:"id"`
	Subject string `json:"name" db:"subject"`
}

//type SubjectVue struct {
//	Subject []string `json:"subject"`
//}
type SubjectVue struct {
	Subject
	Count int `json:"count"`
}
type Chapter struct {
	Id          int    `json:"id" db:"id"`
	Chapter     string `json:"title" db:"chapter"`
	Description string `json:"description" db:"description"`
	Subject     string `json:"typename" db:"subject"`
	Views       string `json:"views" db:"views"`
}

type ChapterVue struct {
	Subject     string `json:"subjectName" `
	Chapter     string `json:"chapterName" `
	Description string `json:"description"`
}

type Exam struct {
	Id            int    `json:"id" db:"id"`
	Subject       string `json:"subject" db:"subject"`
	Chapter       string `json:"chapter" db:"chapter"`
	Question      string `json:"text" db:"question"`
	FirstAnswer   string `json:"first_answer" db:"first_answer"`
	SecondAnswer  string `json:"second_answer" db:"second_answer"`
	ThirdAnswer   string `json:"third_answer" db:"third_answer"`
	FourthAnswer  string `json:"fourth_answer" db:"fourth_answer"`
	CorrectAnswer string `json:"correct_answer" db:"correct_answer"`
}
type ExamVue struct {
	Id            int      `json:"id"` //************************
	Question      string   `json:"text"`
	Answers       []string `json:"answers"`
	CorrectAnswer string   `json:"correctAnswer"`
}
