package dao

import (
	"blog_web/model"
)

type AskBoxDao struct {
	sql []string
}

func NewAskBoxDao() *AskBoxDao {
	return &AskBoxDao{
		sql: []string{
			// 返回已回答的问题
			`SELECT * FROM askbox WHERE is_answered = 1 ORDER BY parent_id DESC, child_id ASC;`,
			// 增加父问题
			`INSERT INTO askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
			VALUES(?, ?, ?, ?, ?, ?)`,
			// 追加子问题
			`INSERT INTO askbox(parent_id, child_id, question, question_time, rainbow)
			VALUES(?, ?, ?, ?, ?)`,
			// 点赞问题
			`UPDATE askbox SET likes = ? WHERE parent_id = ? AND child_id = ?;`,
			// 分页返回所有问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
			 FROM askbox ORDER BY parent_id, child_id DESC LIMIT ?, ? ;`,
			// 分页返回所有未回答问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
			 FROM askbox WHERE is_answered = 0 ORDER BY parent_id, child_id DESC LIMIT ?, ? ;`,
			// 回答问题
			`UPDATE askbox SET answer = ?, answer_time = ?, is_answered =? WHERE parent_id = ? AND child_id = ?;`,
			// 修改问题
			`UPDATE askbox SET answer = ?, answer_time = ? WHERE parent_id = ? AND child_id = ?;`,
			// 删除父问题及底下的子问题
			`DELETE FROM askbox WHERE parent_id = ?;`,
			// 分页返回所有已回答问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent, is_answered
			 FROM askbox WHERE is_answered = 1 ORDER BY parent_id, child_id DESC LIMIT ?, ? ;`,
			// 返回所有已回答父问题数
			`SELECT COUNT(*) FROM askbox WHERE is_answered = 1 AND is_parent = 1 AND child_id = 1;`,
			// 返回所有已回答问题数（按子问题）
			`SELECT COUNT(*) FROM askbox WHERE is_answered = 1;`,
			// 返回所有未回答问题数（按子问题）
			`SELECT COUNT(*) FROM askbox WHERE is_answered = 0;`,
			// 返回所有问题数（按子问题）
			`SELECT COUNT(*) FROM askbox;`,
			// 按父问题序号返回已回答的问题
			`SELECT id, parent_id, child_id, question, question_time, answer, answer_time, rainbow, likes, is_parent FROM askbox WHERE parent_id = ? ;`,
			// 按父问题序号返回父问题下子问题个数
			`SELECT parent_id, COUNT(*) AS child_count FROM askbox GROUP BY parent_id ;`,
			// 返回最大父问题序号
			`SELECT Max(parent_id) from askbox ;`,
		},
	}
}

// // 返回已回答的问题
func (abd *AskBoxDao) GetAnsweredQA() (askboxs []model.Askbox, err error) {
	err = sqldb.Select(&askboxs, abd.sql[0])
	//println("note here")
	//fmt.Println(askboxs)
	return
}

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
	if err != nil {
		println(err.Error())
	}
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

// 分页返回所有已回答问题
func (abd *AskBoxDao) GetAnsweredQAPage(pageStart, PageSize int) (msg []model.Askbox, err error) {
	err = sqldb.Select(&msg, abd.sql[9], pageStart, PageSize)
	return
}

func (abd *AskBoxDao) GetAnsweredParentQuestionCount() (count int, err error) {
	err = sqldb.Get(&count, abd.sql[10])
	return
}

func (abd *AskBoxDao) GetAnsweredQuestionCount() (count int, err error) {
	err = sqldb.Get(&count, abd.sql[11])
	return
}

func (abd *AskBoxDao) GetUnansweredQuestionCount() (count int, err error) {
	err = sqldb.Get(&count, abd.sql[12])
	return
}

func (abd *AskBoxDao) GetAllQuestionCount() (count int, err error) {
	err = sqldb.Get(&count, abd.sql[13])
	return
}

func (abd *AskBoxDao) GetMaxParentQuestionId() (maxID int, err error) {
	err = sqldb.Get(&maxID, "SELECT MAX(parent_id) FROM askbox")
	//println(maxID)
	return
}
