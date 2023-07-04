package dao

import (
	model "blog_web/model"
	"math/rand"
	"time"
)

type SubjectDao struct {
	sql []string
}

type ChapterDao struct {
	sql []string
}
type ExamDao struct {
	sql []string
}

func NewSubjectDao() *SubjectDao {
	return &SubjectDao{
		sql: []string{
			`SELECT id,subject FROM subject ORDER BY id ASC;`,
			`INSERT INTO subject (subject) VALUES (?);`,
			`UPDATE subject SET subject=? WHERE subject=?;`,
			`DELETE FROM subject WHERE subject=?;`,
			`SELECT COUNT(*) FROM chapter WHERE subject=?;`,
		},
	}
}

func NewChapterDao() *ChapterDao {
	return &ChapterDao{
		sql: []string{
			`SELECT * FROM chapter WHERE subject=? ORDER BY id ASC LIMIT ?,?;`,
			`INSERT INTO chapter(subject,chapter,description,views) VALUES (?,?,?,0);`,
			`DELETE FROM chapter WHERE subject=? and chapter=?;`,
			`SELECT * FROM chapter WHERE subject=? ORDER BY id ASC;`,
		},
	}
}

func NewExamDao() *ExamDao {
	return &ExamDao{
		sql: []string{
			`SELECT id,subject,chapter,question,first_answer,second_answer,third_answer,fourth_answer,correct_answer 
FROM exam WHERE subject=? and chapter=? ORDER BY id ASC;`,
			`SELECT COUNT(*) FROM exam WHERE subject=? and chapter=?`,
			`DELETE FROM exam WHERE id=?`,
			`INSERT INTO exam(subject,chapter,question,first_answer,second_answer,third_answer,fourth_answer,correct_answer)
VALUES (?,?,?,?,?,?,?,?)`,
			`UPDATE exam SET subject=?,chapter=?,question=?,first_answer=?,second_answer=?,third_answer=?,
            fourth_answer=?,correct_answer=? WHERE id=?`,
			`UPDATE chapter SET views=views+1 WHERE subject=? and chapter=?`,
			`SELECT id,subject,chapter,question,first_answer,second_answer,third_answer,fourth_answer,correct_answer 
FROM exam WHERE subject=? and chapter=? ORDER BY id ASC LIMIT ?,?;`,
		},
	}
}

//	func (e *SubjectDao) FindAll() (subvuelist model.SubjectVue, err error) {
//		sublist := []model.Subject{}
//		err = sqldb.Select(&sublist, e.sql[0])
//		suball := []string{}
//		for _, sub := range sublist {
//			suball = append(suball, sub.Subject)
//		}
//		subvuelist.Subject = suball
//		return
//	}
func (e *SubjectDao) FindAll() (subvuelist []model.SubjectVue, err error) {
	sublist := []model.Subject{}
	var count int
	err = sqldb.Select(&sublist, e.sql[0])

	for _, sub := range sublist {
		err = sqldb.Get(&count, e.sql[4], sub.Subject)
		subvue := model.SubjectVue{
			Subject: sub,
			Count:   count,
		}

		subvuelist = append(subvuelist, subvue)
	}

	return
}

func (e *SubjectDao) AddSub(subject *model.Subject) error {
	_, err := sqldb.Exec(e.sql[1], subject.Subject)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *SubjectDao) DeleteSub(a string) error {
	_, err := sqldb.Exec(e.sql[3], a)
	if err != nil {
		println(err.Error())
	}
	return err

}

func (e *SubjectDao) UpdateSub(a string, b string) error {
	_, err := sqldb.Exec(e.sql[2], b, a)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *ChapterDao) FindChapter(a string, pageNum int, pageSize int) (chaplist []model.Chapter, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&chaplist, e.sql[0], a, pageStart, pageSize)
	return
}
func (e *ChapterDao) FindAllChapter(a string) (chaplist []model.Chapter, err error) {
	err = sqldb.Select(&chaplist, e.sql[3], a)
	return
}

func (e *ChapterDao) AddChapter(subject string, chapter string, description string) error {
	_, err := sqldb.Exec(e.sql[1], subject, chapter, description)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *ChapterDao) DeleteChapter(subject string, chapter string) error {
	_, err := sqldb.Exec(e.sql[2], subject, chapter)
	if err != nil {
		println(err.Error())
	}
	return err

}

func (e *ExamDao) FindExamLimited(a string, b string) (examvuelist []model.ExamVue, err error) {
	examlist := []model.Exam{}
	var num int
	var count int
	rand.Seed(time.Now().UnixNano())
	count = 0
	err = sqldb.Select(&examlist, e.sql[0], a, b)
	err = sqldb.Get(&num, e.sql[1], a, b)
	_, err = sqldb.Exec(e.sql[5], a, b)
	arr := make([]int, num)
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
	for count < 10 {
		s := rand.Intn(num - 1)
		if arr[s] != 1 {
			examvue := model.ExamVue{
				Id:            examlist[s].Id,
				Question:      examlist[s].Question,
				Answers:       []string{examlist[s].FirstAnswer, examlist[s].SecondAnswer, examlist[s].ThirdAnswer, examlist[s].FourthAnswer},
				CorrectAnswer: examlist[s].CorrectAnswer,
			}
			examvuelist = append(examvuelist, examvue)
			count += 1
			arr[s] = 1
		}
	}
	return
}

func (e *ExamDao) FindExam(subject string, chapter string, pageNum int, pageSize int) (examlist []model.Exam, num int, err error) {
	//examlist := []model.Exam{}
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&examlist, e.sql[6], subject, chapter, pageStart, pageSize)
	err = sqldb.Get(&num, e.sql[1], subject, chapter)
	//for _, exam := range examlist {
	//	examvue := model.Exam{
	//		Id:            exam.Id,
	//		Question:      exam.Question,
	//		Answers:       []string{exam.FirstAnswer, exam.SecondAnswer, exam.ThirdAnswer, exam.FourthAnswer},
	//		CorrectAnswer: exam.CorrectAnswer,
	//	}
	//	examvuelist = append(examvuelist, examvue)
	//
	//}

	return

}
func (e *ExamDao) DeleteExam(id int) error {
	_, err := sqldb.Exec(e.sql[2], id)
	return err
}

func (e *ExamDao) CreateExam(exam *model.Exam) error {
	_, err := sqldb.Exec(e.sql[3], exam.Subject, exam.Chapter, exam.Question, exam.FirstAnswer,
		exam.SecondAnswer, exam.ThirdAnswer, exam.FourthAnswer, exam.CorrectAnswer)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *ExamDao) UpdateExam(exam *model.Exam) error {
	_, err := sqldb.Exec(e.sql[4], exam.Subject, exam.Chapter, exam.Question, exam.FirstAnswer,
		exam.SecondAnswer, exam.ThirdAnswer, exam.FourthAnswer, exam.CorrectAnswer, exam.Id)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *ExamDao) GetMaxId() (maxID int, err error) {
	err = sqldb.Get(&maxID, "SELECT MAX(id) FROM exam")
	//println(maxID)
	return
}
