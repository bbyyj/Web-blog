package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	"testing"
)

// 测试 花式查询
func TestFindElectionComment(t *testing.T) {
	ec := dao.NewElectionCommentDao()

	_, _, err := ec.FindElectionComment("SSE201", 1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find Election Comment Error")
	}
	t.Log("Find Election Comment Success")
}

func TestFindAllElectionComment(t *testing.T) {
	ec := dao.NewElectionCommentDao()

	_, _, err := ec.FindAllElectionComment(1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Election Comment Error")
	}
	t.Log("Find All Election Comment Success")
}

func TestFindElectionCommentByClassification(t *testing.T) {
	ec := dao.NewElectionCommentDao()

	_, _, err := ec.FindElectionCommentByClassification("专必", "计算机组成原理", 1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find Election Comment By Classification Error")
	}
	t.Log("Find Election Comment By Classification Success")
}

// 测试 增删
func TestAddElectionComment(t *testing.T) {
	ec := dao.NewElectionCommentDao()
	var comment model.ElectionComment

	comment.SubjectId = "SSE201"
	comment.SubjectName = "计算机组成原理"
	comment.Classification = "专必"
	comment.Comment = "老师很可爱很可爱"
	err := ec.AddElectionComment(&comment)
	if err != nil {
		println(err.Error())
		t.Fatal("Add Election Comment Error")
	}
	t.Log("Add Election Comment Success")
}

func TestDeleteElectionComment(t *testing.T) {
	ec := dao.NewElectionCommentDao()

	maxId, err := ec.GetMaxId()
	if err != nil {
		println(err.Error())
		t.Fatal("Get MaxId Error")
	}
	t.Log("Get MaxId Success")

	err = ec.DeleteElectionComment(maxId)
	if err != nil {
		println(err.Error())
		t.Fatal("Delete Election Comment Error")
	}
	t.Log("Delete Election Comment Success")
}
