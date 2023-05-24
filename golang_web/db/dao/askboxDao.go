package dao

import "blog_web/model"

type AskBoxDao struct {
	sql []string
}

//func NewAskBoxDao() *AskBoxDao {
//	return &AskBoxDao{
//		sql: []string{
//			// 插入新问题
//			`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
//			VALUES(?, ?, ?, ?, ?, ?)`,
//			// 分页返回所有已回答的问题信息
//			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes,
//			is_parent, is_answered FROM t_askbox WHERE is_answered = 1 LIMIT ?, ?;`,
//			// 返回所有已回答的父问题问答数
//			`SELECT COUNT(*) FROM t_askbox WHERE is_parent = 1 AND is_answered = 1;`,
//			// 分页返回所有问题信息
//			//`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
//			//FROM t_askbox LIMIT ?, ?;`,
//			`SELECT * FROM t_askbox LIMIT ?, ?;`,
//			// 返回所有问题数
//			`SELECT COUNT(*) FROM t_askbox;`,
//			// 返回父问题下的子问题数
//			`SELECT COUNT(*) FROM t_askbox Where parent_id = ? ;`,
//			// 追加回答
//			`UPDATE t_askbox SET answer = ?, answer_time = ?, is_answered =? WHERE parent_id = ? AND child_id = ?;`,
//			// 修改回答
//			`UPDATE t_askbox SET answer = ?, answer_time = ?, is_answered =? WHERE parent_id = ? AND child_id = ?;`,
//			// 删除父问题及底下的子问题
//			`DELETE FROM t_askbox WHERE parent_id = ?;`,
//			// 点赞数+1
//			`UPDATE t_askbox SET likes = ? WHERE parent_id = ? AND child_id = ?;`,
//			// 分页返回所有未回答的问题信息
//			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
//			FROM t_askbox WHERE is_answered = 0 LIMIT ?, ?;`,
//			// 返回未回答的问题数
//			`SELECT COUNT(*) FROM t_askbox Where is_answered = 0 ;`,
//		},
//	}
//}
//
//func (abd *AskBoxDao) AddQuestion(a *model.Askbox) error {
//	//`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
//	_, err := sqldb.Exec(abd.sql[0], a.ParentId, a.ChildId, a.Question, a.QuestionTime, a.Rainbow, a.IsParent)
//	return err
//}
//
//func (abd *AskBoxDao) FindQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[1], pageStart, PageSize)
//	return
//}
//
//func (abd *AskBoxDao) FindAskBoxCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[2])
//	return
//}
//
//func (abd *AskBoxDao) FindAllQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[3], pageStart, PageSize)
//	return
//}
//
//func (abd *AskBoxDao) FindAllAskBoxCount() (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[4])
//	return
//}
//
//func (abd *AskBoxDao) GetChildQuestionCount(a *model.Askbox) (count int, err error) {
//	err = sqldb.Get(&count, abd.sql[5], a.ParentId)
//	return
//}
//
//func (abd *AskBoxDao) AddAnswer(askbox *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[6], askbox.Answer, askbox.AnswerTime, askbox.IsAnswered, askbox.ParentId, askbox.ChildId)
//	return err
//}
//
//func (abd *AskBoxDao) ModifyAnswer(askbox *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[7], askbox.Answer, askbox.AnswerTime, askbox.IsAnswered, askbox.ParentId, askbox.ChildId)
//	return err
//}
//
//func (abd *AskBoxDao) DeleteQuestion(id int) error {
//	_, err := sqldb.Exec(abd.sql[8], id)
//	return err
//}
//
//func (abd *AskBoxDao) ClickLikes(askbox *model.Askbox) error {
//	_, err := sqldb.Exec(abd.sql[7], askbox.Likes, askbox.ParentId, askbox.ChildId)
//	return err
//}
//
//func (abd *AskBoxDao) FindUnansweredQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
//	err = sqldb.Select(&msg, abd.sql[8], pageStart, PageSize)
//	return
//}

func NewAskBoxDao() *AskBoxDao {
	return &AskBoxDao{
		sql: []string{
			// 返回已回答的问题
			`SELECT * FROM t_askbox WHERE is_answered = 1 ORDER BY id ASC;`,
			// 增加父问题
			`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
			VALUES(?, ?, ?, ?, ?, ?)`,
			// 追加子问题
			`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow)
			VALUES(?, ?, ?, ?, ?)`,
			// 点赞问题
			`UPDATE t_askbox SET likes = ? WHERE parent_id = ? AND child_id = ?;`,
			// 分页返回所有问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
			 FROM t_askbox ORDER BY parent_id, child_id ASC LIMIT ?, ? ;`,
			// 分页返回所有未回答问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
			 FROM t_askbox WHERE is_answered = 0 ORDER BY parent_id, child_id ASC LIMIT ?, ? ;`,
			// 回答问题
			`UPDATE t_askbox SET answer = ?, answer_time = ?, is_answered =? WHERE parent_id = ? AND child_id = ?;`,
			// 修改问题
			`UPDATE t_askbox SET answer = ?, answer_time = ? WHERE parent_id = ? AND child_id = ?;`,
			// 删除父问题及底下的子问题
			`DELETE FROM t_askbox WHERE parent_id = ?;`,
		},
	}
}

// 返回已回答的问题
func (abd *AskBoxDao) GetAnsweredQA() (askboxs []model.Askbox, err error) {
	err = sqldb.Select(&askboxs, abd.sql[0])
	return
}

// 增加父问题
func (abd *AskBoxDao) AddNewQuestion(a *model.Askbox) error {
	_, err := sqldb.Exec(abd.sql[1], a.ParentId, a.ChildId, a.Question, a.QuestionTime, a.Rainbow, a.IsParent)
	return err
}

// 追加子问题
func (abd *AskBoxDao) AppendOldQuestion(a *model.Askbox) error {
	_, err := sqldb.Exec(abd.sql[2], a.ParentId, a.ChildId, a.Question, a.QuestionTime, a.Rainbow)
	return err
}

// 点赞问题
func (abd *AskBoxDao) ClickLikes(likes int, parentID int, childID int) error {
	_, err := sqldb.Exec(abd.sql[3], likes, parentID, childID)
	return err
}

// 分页返回所有问题
func (abd *AskBoxDao) GetAllQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
	err = sqldb.Select(&msg, abd.sql[4], pageStart, PageSize)
	return
}

// 分页返回所有未回答问题
func (abd *AskBoxDao) GetUnansweredQA(pageStart, PageSize int) (msg []model.Askbox, err error) {
	err = sqldb.Select(&msg, abd.sql[5], pageStart, PageSize)
	return
}

func (abd *AskBoxDao) AddAnswer(askbox *model.Askbox) error {
	_, err := sqldb.Exec(abd.sql[6], askbox.Answer, askbox.AnswerTime, askbox.IsAnswered, askbox.ParentId, askbox.ChildId)
	return err
}

func (abd *AskBoxDao) ModifyAnswer(askbox *model.Askbox) error {
	_, err := sqldb.Exec(abd.sql[7], askbox.Answer, askbox.AnswerTime, askbox.ParentId, askbox.ChildId)
	return err
}

func (abd *AskBoxDao) DeleteQuestion(id int) error {
	_, err := sqldb.Exec(abd.sql[8], id)
	return err
}
