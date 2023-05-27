package dao

import "blog_web/model"

type ElectionDao struct {
	sql []string
}

func NewElectionDao() *ElectionDao {
	return &ElectionDao{
		sql: []string{
			`SELECT * FROM t_election WHERE subject_id=? ORDER BY subject_id ASC;`,
			`SELECT subject_id,subject_name,teacher FROM t_election WHERE classification=? ORDER BY subject_id ASC LIMIT ?,?;`,
			`SELECT * FROM t_election ORDER BY subject_id ASC LIMIT ?,?;`,
			`DELETE FROM t_election WHERE subject_id=?`,
			`INSERT INTO t_election(subject_id,subject_name,teacher,classification,credit,college,attendance,score,evaluation)
VALUES (?,?,?,?,?,?,?,?,?)`,
			`UPDATE t_election SET subject_name=?,teacher=?,classification=?,credit=?,college=?,attendance=?,
            score=?,evaluation=? WHERE subject_id=?`,
		},
	}
}

func (e *ElectionDao) FindDetailedElection(subjectId string) (electionlist []model.ElectionDetailed, err error) {
	err = sqldb.Select(&electionlist, e.sql[0], subjectId)
	return
}
func (e *ElectionDao) FindAllElection(classification string, pageNum int, pageSize int) (electionlist []model.Election, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[1], classification, pageStart, pageSize)
	return
}
func (e *ElectionDao) FindAllDetailedElection(pageNum int, pageSize int) (electionlist []model.ElectionDetailed, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionlist, e.sql[2], pageStart, pageSize)
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
