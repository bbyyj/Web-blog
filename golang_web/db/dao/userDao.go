package dao

import "blog_web/model"

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM user WHERE username = ? AND password = ?;`,
			`SELECT nickname, avatar FROM user Limit 1`,
		},
	}
}

func (u *UserDao) FindUserByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[0], username, password)
	return &user, err
}

func (u *UserDao) FindInfo() (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[1])
	return &user, err
}
