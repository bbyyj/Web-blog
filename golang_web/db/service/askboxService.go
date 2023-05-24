package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
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

func (a *AskBoxService) AddNewQuestion(askbox *model.Askbox) error {
	return a.askboxDao.AddNewQuestion(askbox)
}

func (a *AskBoxService) AppendOldQuestion(askbox *model.Askbox) error {
	return a.askboxDao.AddNewQuestion(askbox)
}

func (a *AskBoxService) ClickLikes(likes int, parentID int, childID int) error {
	return a.askboxDao.ClickLikes(likes, parentID, childID)
}

func (a *AskBoxService) GetAllQA(pageNum, pageSize int) ([]model.Askbox, int, error) {
	pageStart := (pageNum - 1) * pageSize
	messages, err := a.askboxDao.GetAllQA(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("Get All QA error:%v", err)
		return nil, 0, err
	}
	count, _ := a.askboxDao.GetAllQuestionCount()
	return messages, count, nil
}

func (a *AskBoxService) GetUnansweredQA(pageNum, pageSize int) ([]model.Askbox, int, error) {
	pageStart := (pageNum - 1) * pageSize
	messages, err := a.askboxDao.GetUnansweredQA(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("Get Unanswered QA error:%v", err)
		return nil, 0, err
	}
	count, _ := a.askboxDao.GetUnansweredQuestionCount()
	return messages, count, nil
}

func (a *AskBoxService) GetAnsweredQAPage(pageNum, pageSize int) ([]model.Askbox, int, error) {
	pageStart := (pageNum - 1) * pageSize
	messages, err := a.askboxDao.GetAnsweredQAPage(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("Get Answered QA error:%v", err)
		return nil, 0, err
	}
	count, _ := a.askboxDao.GetAnsweredQuestionCount()
	return messages, count, nil
}

func (a *AskBoxService) AddAnswer(askbox *model.Askbox) error {
	return a.askboxDao.AddAnswer(askbox)
}

func (a *AskBoxService) ModifyAnswer(askbox *model.Askbox) error {
	return a.askboxDao.ModifyAnswer(askbox)
}

func (a *AskBoxService) DeleteQuestion(id int) error {
	return a.askboxDao.DeleteQuestion(id)
}
