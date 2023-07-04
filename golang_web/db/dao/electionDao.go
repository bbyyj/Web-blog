package dao

import "blog_web/model"

type ElectionDao struct {
	sql []string
}

func NewElectionDao() *ElectionDao {
	return &ElectionDao{
		sql: []string{
			`SELECT * FROM election WHERE subject_id=? ORDER BY subject_id ASC;`,
			`SELECT subject_id,subject_name,teacher FROM election WHERE classification=? ORDER BY subject_id ASC LIMIT ?,?;`,
			`SELECT * FROM election WHERE classification=? ORDER BY subject_id ASC LIMIT ?,?;`,
			`DELETE FROM election WHERE subject_id=?`,
			`INSERT INTO election(subject_id,subject_name,teacher,classification,credit,college,attendance,score,evaluation)
VALUES (?,?,?,?,?,?,?,?,?)`,
			`UPDATE election SET subject_name=?,teacher=?,classification=?,credit=?,college=?,attendance=?,
            score=?,evaluation=? WHERE subject_id=?`,
			`SELECT * FROM election ORDER BY subject_id ASC LIMIT ?,?;`,
			`SELECT COUNT(*) FROM election WHERE classification=?;`,
			`SELECT COUNT(*) FROM election ;`,
			`SELECT subject_id,subject_name,teacher FROM election WHERE classification=? ORDER BY subject_id ASC;`,
			`SELECT subject_id,subject_name,teacher FROM election ORDER BY subject_id ASC LIMIT ?,?;`,
		},
	}
}

func (e *ElectionDao) FindDetailedElection(subjectId string) (electionlist []model.ElectionDetailed, err error) {
	err = sqldb.Select(&electionlist, e.sql[0], subjectId)
	return
}
func (e *ElectionDao) FindAllElection(classification string, pageNum int, pageSize int) (electionlist []model.Election, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[1], classification, pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[7], classification)
	return
}

func (e *ElectionDao) FindAllElectionNoClass(pageNum int, pageSize int) (electionlist []model.Election, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[10], pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[8])
	return
}
func (e *ElectionDao) FindAllElectionNopage(classification string) (electionlist []model.Election, err error) {
	err = sqldb.Select(&electionlist, e.sql[9], classification)
	return
}

func (e *ElectionDao) FindElectionByClass(classification string, pageNum int, pageSize int) (electionlist []model.ElectionDetailed, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[2], classification, pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[7], classification)

	return
}
func (e *ElectionDao) FindAllDetailedElection(pageNum int, pageSize int) (electionlist []model.ElectionDetailed, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[6], pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[8])
	return
}
func (e *ElectionDao) DeleteElection(subjectId string) error {
	_, err := sqldb.Exec(e.sql[3], subjectId)
	if err != nil {
		println(err.Error())
	}
	return err
}
func (e *ElectionDao) CreateElection(election *model.ElectionDetailed) error {
	_, err := sqldb.Exec(e.sql[4], election.SubjectId, election.SubjectName, election.Teacher, election.Classification,
		election.Credit, election.College, election.Attendance, election.Score, election.Evaluation)
	if err != nil {
		println(err.Error())
	}
	return err
}
func (e *ElectionDao) UpdateElection(election *model.ElectionDetailed) error {
	_, err := sqldb.Exec(e.sql[5], election.SubjectName, election.Teacher, election.Classification,
		election.Credit, election.College, election.Attendance, election.Score, election.Evaluation, election.SubjectId)
	if err != nil {
		println(err.Error())
	}
	return err
}
