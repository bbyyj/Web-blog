package dao

import "blog_web/model"

type TacitDao struct {
	sql []string
}

func NewTacitDao() *TacitDao {
	return &TacitDao{
		sql: []string{
			`SELECT id, question,first_answer,second_answer,third_answer,fourth_answer,correct_answer FROM t_tacit ORDER BY id ASC;`,
			//`SELECT id, Question,FirstAnswer,SecondAnswer,ThirdAnswer,FourthAnswer FROM t_tacit ORDER BY id DESC LIMIT ?, ?;`,
			`SELECT COUNT(*) FROM t_tacit`,
			`INSERT INTO t_tacit ( question,first_answer,second_answer,third_answer,fourth_answer,,correct_answer) VALUES (?,?,?,?,?,?);`,
			`DELETE FROM t_tacit WHERE id = ?;`,
			`UPDATE t_tacit SET question=?,first_answer=?,second_answer=?,third_answer=?,fourth_answer=?,,correct_answer=? WHERE id = ?;`,
		},
	}
}

func (e *TacitDao) FindAll() (tacitvueList []model.TacitVue, err error) {
	tacitList := []model.Tacit{}
	err = sqldb.Select(&tacitList, e.sql[0])
	for _, tacit := range tacitList {
		tacitvue := model.TacitVue{
			Question:      tacit.Question,
			Answers:       []string{tacit.FirstAnswer, tacit.SecondAnswer, tacit.ThirdAnswer, tacit.FourthAnswer},
			CorrectAnswer: tacit.CorrectAnswer,
		}
		tacitvueList = append(tacitvueList, tacitvue)
	}
	return
}

//func (e *EssayDao) FindLimited(pageStart, pageSize int) (essays []model.Essay, err error) {
//	err = sqldb.Select(&essays, e.sql[1], pageStart, pageSize)
//	return
//}
//
//func (e *EssayDao) FindTotalCount() (count int, err error) {
//	err = sqldb.Get(&count, e.sql[2])
//	return
//}
//
//func (e *EssayDao) Add(essay *model.Essay) error {
//	_, err := sqldb.Exec(e.sql[3], essay.CreateTime, essay.Content)
//	return err
//}
//
//func (e *EssayDao) Delete(id int) error {
//	_, err := sqldb.Exec(e.sql[4], id)
//	return err
//}
//
//func (e *EssayDao) Update(essay *model.Essay) error {
//	_, err := sqldb.Exec(e.sql[5], essay.Content, essay.Id)
//	return err
//}
