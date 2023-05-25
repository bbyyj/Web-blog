package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
	"fmt"
)

type AskBoxService struct {
	askboxDao *dao.AskBoxDao
}

func NewAskBoxService() *AskBoxService {
	return &AskBoxService{
		askboxDao: dao.NewAskBoxDao(),
	}
}

func groupByParentID(askboxs []model.Askbox) [][]model.Askbox {
	groups := make(map[int][]model.Askbox)
	for _, ab := range askboxs {
		groups[ab.ParentId] = append(groups[ab.ParentId], ab)
	}

	var result [][]model.Askbox
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func (a *AskBoxService) GetAnsweredQA() ([][]model.Askbox, int, error) {

	askboxs, err := a.askboxDao.GetAnsweredQA()
	if err != nil {
		fmt.Println(askboxs)
		fmt.Printf("askboxs的类型是：%T\n", askboxs)
	}

	result := groupByParentID(askboxs)
	fmt.Println(result)

	count, _ := a.askboxDao.GetAnsweredParentQuestionCount()

	return result, count, nil
}

func (a *AskBoxService) AddNewQuestion(askbox *model.Askbox) error {
	return a.askboxDao.AddNewQuestion(askbox)
}

func (a *AskBoxService) GetMaxParentQuestionId() (int, error) {
	maxParentQuestionId, err := a.askboxDao.GetMaxParentQuestionId()
	if err != nil {
		utils.Logger().Warning("Get Max ParentQuestionId error:%v", err)
		return 0, err
	}
	return maxParentQuestionId, nil
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
