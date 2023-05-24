package service

import (
	"blog_web/db/dao"
)

type AskBoxService struct {
	askboxDao *dao.AskBoxDao
}

func NewAskBoxService() *AskBoxService {
	return &AskBoxService{
		askboxDao: dao.NewAskBoxDao(),
	}
}

// 报错？
//func (a *AskBoxService) GetAnsweredQA([]model.Askbox, error) {
//	return a.askboxDao.GetAnsweredQA()
//}
