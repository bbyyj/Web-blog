package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type TacitService struct {
	tacitDao *dao.TacitDao
}

func NewTacitService() *TacitService {
	return &TacitService{
		tacitDao: dao.NewTacitDao(),
	}
}

func (e *TacitService) GetAll() ([]model.TacitVue, error) {

	return e.tacitDao.FindAll()
}

//func (e *EssayService) GetLimited(pageNum, pageSize int) ([]model.Essay, error) {
//	pageStart := (pageNum - 1) * pageSize
//	return e.essayDao.FindLimited(pageStart, pageSize)
//}
//
//func (e *EssayService) GetCount() (int, error) {
//	return e.essayDao.FindTotalCount()
//}
//
//func (e *EssayService) AddEssay(essay *model.Essay) error {
//	return e.essayDao.Add(essay)
//}
//
//func (e *EssayService) DeleteEssay(id int) error {
//	return e.essayDao.Delete(id)
//}
//
//func (e *EssayService) UpdateEssay(essay *model.Essay) error {
//	return e.essayDao.Update(essay)
//}
