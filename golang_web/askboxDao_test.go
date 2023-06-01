package main

import (
	"blog_web/db/dao"
	"testing"
)

// 测试一系列GetCount函数
func TestGetAnsweredParentQuestionCount(t *testing.T) {
	abd := dao.NewAskBoxDao()
	count, err := abd.GetAnsweredParentQuestionCount()
	if err != nil {
		t.Fatal("Get Answered Parent Question Count Error")
	}
	t.Logf("Get Answered Parent Question Count Success, count = %v", count)
}

func TestGetAnsweredQuestionCount(t *testing.T) {
	abd := dao.NewAskBoxDao()
	count, err := abd.GetAnsweredQuestionCount()
	if err != nil {
		t.Fatal("GetAnsweredQuestionCount Error")
	}
	t.Logf("GetAnsweredQuestionCount Success, count = %v", count)
}

func TestGetUnansweredQuestionCount(t *testing.T) {
	abd := dao.NewAskBoxDao()
	count, err := abd.GetUnansweredQuestionCount()
	if err != nil {
		t.Fatal("Get Unanswered Question Count Error")
	}
	t.Logf("Get Unanswered Question Count Success, count = %v", count)
}

func TestGetAllQuestionCountCount(t *testing.T) {
	abd := dao.NewAskBoxDao()
	count, err := abd.GetAllQuestionCount()
	if err != nil {
		t.Fatal("Get All Question Count Error")
	}
	t.Logf("Get All Question Count Success, count = %v", count)
}

// 测试一系列GetQA函数

//func (abd *AskBoxDao) GetAnsweredQAPage(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[9], pageStart, PageSize)
//	return
//}
//

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

//func TestGetAnsweredQAPage(t *testing.T) {
//	abd := dao.NewAskBoxDao()
//	err := abd.GetAnsweredQAPage(1,10)
//	if err != nil {
//		t.Fatal("Get Answered QAPage Error")
//	}
//	t.Logf("Get Answered QAPage Success, QAPage = %v", count)
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
