package model

import "time"

type Askbox struct {
	Id           int       `db:"id" json:"id"`
	parentId     int       `db:"parent_id" json:"parent_id"`
	childId      int       `db:"child_id" json:"child_id"`
	Answer       string    `db:"answer" json:"answer"`
	AnswerTime   time.Time `db:"answer_time" json:"answer_Time"`
	Question     string    `db:"question" json:"question"`
	QuestionTime time.Time `db:"question_time" json:"question_Time"`
}
