package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	_ "blog_web/model"
	_ "math/rand"
	"testing"
	_ "time"
)

// 测试 增删改查 subject
//func TestFindAll(t *testing.T) {
//	e := dao.NewSubjectDao()
//
//	_, err := e.FindAll()
//	if err != nil {
//		println(err.Error())
//		t.Fatal("Find All Error")
//	}
//	t.Log("Find All Success")
//}

func TestAddSub(t *testing.T) {
	e := dao.NewSubjectDao()

	var s model.Subject
	s.Subject = "你没见过的全新科目"

	err := e.AddSub(&s)
	if err != nil {
		println(err.Error())
		t.Fatal("Add Sub Error")
	}
	t.Log("Add Sub Success")
}

func TestUpdateSub(t *testing.T) {
	e := dao.NewSubjectDao()

	err := e.UpdateSub("你没见过的全新科目", "你没见过的全新科目2")
	if err != nil {
		println(err.Error())
		t.Fatal("Update Sub Error")
	}
	t.Log("Update Sub Success")
}

func TestDeleteSub(t *testing.T) {
	e := dao.NewSubjectDao()

	err := e.DeleteSub("你没见过的全新科目2")
	if err != nil {
		println(err.Error())
		t.Fatal("Delete Sub Error")
	}
	t.Log("Delete Sub Success")
}

// 测试 增删查 chapter
func TestFindChapter(t *testing.T) {
	e := dao.NewChapterDao()

	_, err := e.FindChapter("总线", 1, 10)
	if err != nil {
		println(err.Error())
		t.Fatal("Find Chapter Error")
	}
	t.Log("Find Chapter Success")
}

func TestAllChapter(t *testing.T) {
	e := dao.NewChapterDao()

	_, err := e.FindAllChapter("总线")
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Chapter Error")
	}
	t.Log("Find All Chapter Success")
}

func TestAddChapter(t *testing.T) {
	e := dao.NewChapterDao()

	err := e.AddChapter("测试完了没有", "测试完了没有", "测试完了没有")
	if err != nil {
		println(err.Error())
		t.Fatal("Add Chapter Error")
	}
	t.Log("Add Chapter Success")
}

func TestDeleteChapter(t *testing.T) {
	e := dao.NewChapterDao()

	err := e.DeleteChapter("测试完了没有", "测试完了没有")
	if err != nil {
		println(err.Error())
		t.Fatal("Delete Chapter Error")
	}
	t.Log("Delete Chapter Success")
}

// 测试 查 exam
func TestFindExamLimited(t *testing.T) {
	e := dao.NewExamDao()

	_, err := e.FindExamLimited("计算机组成原理", "总线")
	if err != nil {
		println(err.Error())
		t.Fatal("Find Exam Limited Error")
	}
	t.Log("Find Exam Limited Success")
}

func TestFindExam(t *testing.T) {
	e := dao.NewExamDao()

	_, _, err := e.FindExam("计算机组成原理", "总线", 1, 10)
	if err != nil {
		println(err.Error())
		t.Fatal("Find Exam Error")
	}
	t.Log("Find Exam Success")
}

// 测试 增改删 exam
func TestCreateExam(t *testing.T) {
	e := dao.NewExamDao()

	var exam model.Exam
	exam.Subject = "计算机组成原理"
	exam.Chapter = "总线"
	exam.Question = "q"
	exam.FirstAnswer = "1"
	exam.SecondAnswer = "2"
	exam.ThirdAnswer = "3"
	exam.FourthAnswer = "4"
	exam.CorrectAnswer = "1"

	err := e.CreateExam(&exam)
	if err != nil {
		println(err.Error())
		t.Fatal("Create Exam Error")
	}
	t.Log("Create Exam Success")
}

func TestUpdateExam(t *testing.T) {
	e := dao.NewExamDao()

	var exam model.Exam
	exam.Subject = "计算机组成原理"
	exam.Chapter = "总线"
	exam.Question = "qq"
	exam.FirstAnswer = "11"
	exam.SecondAnswer = "22"
	exam.ThirdAnswer = "33"
	exam.FourthAnswer = "44"
	exam.CorrectAnswer = "11"
	maxId, err := e.GetMaxId()
	exam.Id = maxId
	if err != nil {
		println(err.Error())
		t.Fatal("Get MaxId Error")
	}
	t.Log("Get MaxId Success")

	err = e.UpdateExam(&exam)
	if err != nil {
		println(err.Error())
		t.Fatal("Update Exam Error")
	}
	t.Log("Update Exam Success")
}

func TestDeleteExam(t *testing.T) {
	e := dao.NewExamDao()

	maxId, err := e.GetMaxId()
	if err != nil {
		println(err.Error())
		t.Fatal("Get MaxId Error")
	}
	t.Log("Get MaxId Success")

	err = e.DeleteExam(maxId)
	if err != nil {
		println(err.Error())
		t.Fatal("Update Exam Error")
	}
	t.Log("Update Exam Success")
}
