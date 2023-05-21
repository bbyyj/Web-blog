package model

type Tacit struct {
	Id            int    `json:"id" db:"id"`
	Question      string `json:"text" db:"question"`
	FirstAnswer   string `json:"first_answer" db:"first_answer"`
	SecondAnswer  string `json:"second_answer" db:"second_answer"`
	ThirdAnswer   string `json:"third_answer" db:"third_answer"`
	FourthAnswer  string `json:"fourth_answer" db:"fourth_answer"`
	CorrectAnswer string `json:"correct_answer" db:"correct_answer"`
}
type TacitVue struct {
	Question      string   `json:"text"`
	Answers       []string `json:"answers"`
	CorrectAnswer string   `json:"correctAnswer"`
}
