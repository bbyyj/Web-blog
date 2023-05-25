package dao

import model "blog_web/model"

type SubjectDao struct {
	sql []string
}

type ChapterDao struct {
	sql []string
}

func NewSubjectDao() *SubjectDao {
	return &SubjectDao{
		sql: []string{
			`SELECT subject FROM t_subject ORDER BY id ASC;`,
			//`SELECT id, Question,FirstAnswer,SecondAnswer,ThirdAnswer,FourthAnswer FROM t_tacit ORDER BY id DESC LIMIT ?, ?;`,
			//`SELECT COUNT(*) FROM t_tacit`,
			//`INSERT INTO t_tacit ( question,first_answer,second_answer,third_answer,fourth_answer,,correct_answer) VALUES (?,?,?,?,?,?);`,
			//`DELETE FROM t_tacit WHERE id = ?;`,
			//`UPDATE t_tacit SET question=?,first_answer=?,second_answer=?,third_answer=?,fourth_answer=?,,correct_answer=? WHERE id = ?;`,
		},
	}
}

func NewChapterDao() *ChapterDao {
	return &ChapterDao{
		sql: []string{
			`SELECT id,subject,chapter FROM t_chapter WHERE subject=? ORDER BY id ASC;`,
		},
	}
}

func (e *SubjectDao) FindAll() (subvuelist model.SubjectVue, err error) {
	sublist := []model.Subject{}
	err = sqldb.Select(&sublist, e.sql[0])
	suball := []string{}
	for _, sub := range sublist {
		suball = append(suball, sub.Subject)
	}
	subvuelist.Subject = suball
	return
}

func (e *ChapterDao) FindChapter(a string) (chapvuelist model.ChapterVue, err error) {
	chaplist := []model.Chapter{}
	err = sqldb.Select(&chaplist, e.sql[0], a)
	chapall := []string{}
	for _, chap := range chaplist {
		chapall = append(chapall, chap.Chapter)
	}
	chapvuelist.Subject = a
	chapvuelist.Chapter = chapall
	return
}
