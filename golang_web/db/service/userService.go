package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type UserService struct {
	userDao *dao.UserDao
	tagDao  *dao.TagDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
		tagDao:  dao.NewTagDao(),
	}
}

func (u *UserService) CheckUser(username, password string) (*model.User, error) {
	println("note here serve")
	return u.userDao.FindUserByUsernameAndPassword(username, password)
}

func (u *UserService) GetInfo() (*model.User, int, error) {
	info, err := u.userDao.FindInfo()
	if err != nil {
		return nil, 0, err
	}
	count, _ := u.tagDao.GetTagsCount()

	return info, count, nil
}
