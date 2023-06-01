package dao

import "blog_web/model"

/*
* @Author: mgh
* @Date: 2022/3/2 11:30
* @Desc:
 */

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM t_user WHERE username = ? AND password = ?;`,
			`SELECT nickname, email, avatar, github, csdn FROM t_user Limit 1`,
		},
	}
}

func (u *UserDao) FindUserByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	println("note here dao 1")
	println(username)
	println(password)
	err := sqldb.Get(&user, u.sql[0], username, password)
	println("note here dao 2")
	println(user.Username)
	println(user.Password)
	// username admin
	// password 123456
	// password coded e10adc3949ba59abbe56e057f20f883e
	return &user, err
}

func (u *UserDao) FindInfo() (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[1])
	return &user, err
}
