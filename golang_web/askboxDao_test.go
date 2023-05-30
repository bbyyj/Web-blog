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
