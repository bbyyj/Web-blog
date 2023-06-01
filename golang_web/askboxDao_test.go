package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	_ "blog_web/model"
	"testing"
	"time"
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
func TestGetAnsweredQA(t *testing.T) {
	abd := dao.NewAskBoxDao()
	qa, err := abd.GetAnsweredQA()
	if err != nil {
		t.Fatal("Get Answered QA Error")
	}
	t.Logf("Get Answered QA Success, QAPage = %v", qa)
}

func TestGetAnsweredQAPage(t *testing.T) {
	abd := dao.NewAskBoxDao()
	qaPage, err := abd.GetAnsweredQAPage(1, 10)
	if err != nil {
		t.Fatal("Get Answered QAPage Error")
	}
	t.Logf("Get Answered QAPage Success, QAPage = %v", qaPage)
}

func TestGetAllQA(t *testing.T) {
	abd := dao.NewAskBoxDao()
	qaPage, err := abd.GetAnsweredQA()
	if err != nil {
		t.Fatal("Get All QA Error")
	}
	t.Logf("Get All QA Success, QAPage = %v", qaPage)
}

func TestGetUnansweredQA(t *testing.T) {
	abd := dao.NewAskBoxDao()
	qaPage, err := abd.GetUnansweredQA(1, 10)
	if err != nil {
		t.Fatal("Get Unanswered QA Error")
	}
	t.Logf("Get Unanswered QA Success, QAPage = %v", qaPage)
}

// 测试其他Get函数
func TestGetMaxParentQuestionId(t *testing.T) {
	abd := dao.NewAskBoxDao()
	maxId, err := abd.GetMaxParentQuestionId()
	if err != nil {
		t.Fatal("Get Max Parent Question Id Error")
	}
	t.Logf("Get Max Parent Question Id Success, maxId = %v", maxId)
}

// 测试POST-PUT-DELETE函数
// 增加父问题，点赞父问题，回答父问题，修改父问题回答，增加子问题，删除整个问题
func TestAddNewQuestion(t *testing.T) {
	abd := dao.NewAskBoxDao()
	var a model.Askbox
	a.ParentId = 1000
	a.ChildId = 1
	a.Question = "go test parent question"
	a.QuestionTime = time.Now()
	a.IsParent = true
	a.Rainbow = true
	err := abd.AddNewQuestion(&a)
	if err != nil {
		t.Fatal("Add New Question Error")
	}
	t.Log("Add New Question Success")
}

func TestClickLikes(t *testing.T) {
	abd := dao.NewAskBoxDao()
	err := abd.ClickLikes(0, 1000, 1)
	if err != nil {
		t.Fatal("Click Likes Error")
	}
	t.Log("Click Likes Success")
}

func TestAddAnswer(t *testing.T) {
	abd := dao.NewAskBoxDao()
	var a model.Askbox
	a.ParentId = 1000
	a.ChildId = 1
	a.Answer = "go test parent answer"
	a.AnswerTime = time.Now()
	a.IsAnswered = true
	err := abd.AddAnswer(&a)
	if err != nil {
		t.Fatal("Add Answer Error")
	}
	t.Log("Add Answer Success")
}

func TestAppendOldQuestion(t *testing.T) {
	abd := dao.NewAskBoxDao()
	var a model.Askbox
	a.ParentId = 1000
	a.ChildId = 1
	a.Question = "go test child question"
	a.QuestionTime = time.Now()
	a.Rainbow = true
	err := abd.AppendOldQuestion(&a)
	if err != nil {
		t.Fatal("Append Old Question Error")
	}
	t.Log("Append Old Question Success")
}

func TestModifyAnswer(t *testing.T) {
	abd := dao.NewAskBoxDao()
	var a model.Askbox
	a.ParentId = 1000
	a.ChildId = 1
	a.Answer = "go test modify answer"
	a.AnswerTime = time.Now()
	err := abd.ModifyAnswer(&a)
	if err != nil {
		t.Fatal("Modify Answer Error")
	}
	t.Log("Modify Answer Success")
}

func TestDeleteQuestion(t *testing.T) {
	abd := dao.NewAskBoxDao()
	err := abd.DeleteQuestion(1000)
	if err != nil {
		t.Fatal("Delete Question Error")
	}
	t.Log("Delete Question Success")
}
