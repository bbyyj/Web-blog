package main

import (
	"blog_web/db/dao"
	"testing"
)

func TestGetAnsweredQuestionCount(t *testing.T) {
	abd := dao.NewAskBoxDao()
	count, err := abd.GetAnsweredQuestionCount()
	if err != nil {
		t.Fatal("GetAnsweredQuestionCount Error")
	}
	t.Logf("GetAnsweredQuestionCount Success, count = %v", count)
}

//
//func (abd *AskBoxDao) GetAnsweredQAPage(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[9], pageStart, PageSize)
//	return
//}
//
//func (abd *AskBoxDao) GetAnsweredParentQuestionCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[10])
//	return
//}
//
//func (abd *AskBoxDao) GetAnsweredQuestionCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[11])
//	return
//}
//
//func (abd *AskBoxDao) GetUnansweredQuestionCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[12])
//	return
//}
//
//func (abd *AskBoxDao) GetAllQuestionCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[13])
//	return
//}
//
//func (abd *AskBoxDao) GetMaxParentQuestionId() (maxID int, err error) {
//	err = sqldb.Get(&maxID, "SELECT MAX(parent_id) FROM t_askbox")
//	println(maxID)
//	return
//}
//
//func (abd *AskBoxDao) GetUnansweredQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[5], pageStart, PageSize)
//	return
//}
//
//func (abd *AskBoxDao) GetAllQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[4], pageStart, PageSize)
//	if err != nil {
//		println(err.Error())
//	}
//	return
//}
//
//func (abd *AskBoxDao) GetAnsweredQA() (askboxs []model.Askbox, err error) {
//	// abd.sql[0]: `SELECT * FROM t_askbox WHERE is_answered = 1 ORDER BY parent_id, child_id ASC;`,
//	err = sqldb.Select(&askboxs, abd.sql[0])
//	return
//}

//////////////////////////////////////////
//func (abd *AskBoxDao) AddNewQuestion(a *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[1], a.ParentId, a.ChildId, a.Question, a.QuestionTime, a.Rainbow, a.IsParent)
//	return err
//}
//
//// 追加子问题
//func (abd *AskBoxDao) AppendOldQuestion(a *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[2], a.ParentId, a.ChildId, a.Question, a.QuestionTime, a.Rainbow)
//	return err
//}
//
//// 点赞问题
//func (abd *AskBoxDao) ClickLikes(likes int, parentID int, childID int) error {
//	_, err := sqldb.Exec(abd.sql[3], likes, parentID, childID)
//	return err
//}
//
//func (abd *AskBoxDao) AddAnswer(askbox *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[6], askbox.Answer, askbox.AnswerTime, askbox.IsAnswered, askbox.ParentId, askbox.ChildId)
//	return err
//}
//
//func (abd *AskBoxDao) ModifyAnswer(askbox *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[7], askbox.Answer, askbox.AnswerTime, askbox.ParentId, askbox.ChildId)
//	return err
//}
//
//func (abd *AskBoxDao) DeleteQuestion(id int) error {
//	_, err := sqldb.Exec(abd.sql[8], id)
//	return err
//}
