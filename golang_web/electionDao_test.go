package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	"testing"
)

// 测试 花式查询
func TestFindAllElection(t *testing.T) {
	election := dao.NewElectionDao()

	_, _, err := election.FindAllElection("专必", 1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Election Error")
	}
	t.Log("Find All Election Success")
}

func TestFindAllElectionNoClass(t *testing.T) {
	election := dao.NewElectionDao()

	_, _, err := election.FindAllElectionNoClass(1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Election NoClass Error")
	}
	t.Log("Find All Election NoClass Success")
}

func TestFindAllElectionNopage(t *testing.T) {
	election := dao.NewElectionDao()

	_, err := election.FindAllElectionNopage("专必")
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Election Nopage Error")
	}
	t.Log("Find All Election Nopage Success")
}

func TestFindAllDetailedElection(t *testing.T) {
	election := dao.NewElectionDao()

	_, _, err := election.FindAllDetailedElection(1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Detailed Election Error")
	}
	t.Log("Find All Detailed Election Success")
}

func TestFindElectionByClass(t *testing.T) {
	election := dao.NewElectionDao()

	_, _, err := election.FindElectionByClass("专必", 1, 5)
	if err != nil {
		println(err.Error())
		t.Fatal("Find Election By Class Error")
	}
	t.Log("Find Election By Class Success")
}

func TestFindDetailedElection(t *testing.T) {
	election := dao.NewElectionDao()

	_, err := election.FindDetailedElection("SSE201")
	if err != nil {
		println(err.Error())
		t.Fatal("Find Detailed Election Error")
	}
	t.Log("Find Detailed Election Success")
}

// 测试 增改删
func TestCreateElection(t *testing.T) {
	election := dao.NewElectionDao()

	var ele model.ElectionDetailed
	ele.SubjectId = "SSE99"
	ele.SubjectName = "很棒的一门课"
	ele.Teacher = "很棒的一位老师"
	ele.Classification = "专必"
	ele.Credit = "99"
	ele.College = "软件工程学院"
	ele.Attendance = "无考勤"
	ele.Score = "99"
	ele.Evaluation = "无考试"

	err := election.CreateElection(&ele)
	if err != nil {
		println(err.Error())
		t.Fatal("Create Election Error")
	}
	t.Log("Create Election Success")
}

func TestUpdateElection(t *testing.T) {
	election := dao.NewElectionDao()

	var ele model.ElectionDetailed
	ele.SubjectId = "SSE99"
	ele.SubjectName = "很棒棒的一门课！"
	ele.Teacher = "很棒棒的一位老师！"
	ele.Classification = "专必！"
	ele.Credit = "99"
	ele.College = "软件工程学院！"
	ele.Attendance = "无考勤！"
	ele.Score = "999"
	ele.Evaluation = "无考试！"

	err := election.UpdateElection(&ele)
	if err != nil {
		println(err.Error())
		t.Fatal("Update Election Error")
	}
	t.Log("Update Election Success")
}

func TestDeleteElection(t *testing.T) {
	election := dao.NewElectionDao()

	err := election.DeleteElection("SSE99")
	if err != nil {
		println(err.Error())
		t.Fatal("Delete Election Error")
	}
	t.Log("Delete Election Success")
}
