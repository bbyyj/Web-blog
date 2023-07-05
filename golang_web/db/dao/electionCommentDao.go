package dao

import "blog_web/model"

type ElectionCommentDao struct {
	sql []string
}

func NewElectionCommentDao() *ElectionCommentDao {
	return &ElectionCommentDao{
		sql: []string{
			`SELECT * FROM election_comment WHERE subject_id=? ORDER BY id ASC LIMIT ?,?;`,
			`SELECT * FROM election_comment ORDER BY id ASC LIMIT ?,?;`,
			`DELETE FROM election_comment WHERE id=?`,
			`INSERT INTO election_comment(subject_id,subject_name,classification,comment) VALUES (?,?,?,?)`,
			`SELECT COUNT(*) FROM election_comment `,
			`SELECT COUNT(*) FROM election_comment WHERE subject_id=?`,
			`SELECT subject_name FROM t_election WHERE subject_id=?`,
			`SELECT * FROM election_comment WHERE classification=? and subject_name=? ORDER BY id ASC LIMIT ?,?;`,
			`SELECT COUNT(*) FROM election_comment WHERE classification=?`,
			`SELECT classification FROM t_election WHERE subject_id=?`,
			`SELECT COUNT(*) FROM election_comment WHERE classification=? and subject_name=?`,
		},
	}
}
func (e *ElectionCommentDao) FindElectionComment(subjectId string, pageNum int, pageSize int) (electionCommentList []model.ElectionComment, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionCommentList, e.sql[0], subjectId, pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[5], subjectId)
	return
}

func (e *ElectionCommentDao) FindAllElectionComment(pageNum int, pageSize int) (electionCommentList []model.ElectionComment, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionCommentList, e.sql[1], pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[4])
	if err != nil {
		println(err.Error())
	}
	return
}
func (e *ElectionCommentDao) FindElectionCommentByClassification(classification string, subjectName string, pageNum int, pageSize int) (electionCommentList []model.ElectionComment, count int, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionCommentList, e.sql[7], classification, subjectName, pageStart, pageSize)
	err = sqldb.Get(&count, e.sql[10], classification, subjectName)
	if err != nil {
		println(err.Error())
	}
	return
}
func (e *ElectionCommentDao) DeleteElectionComment(id int) error {
	_, err := sqldb.Exec(e.sql[2], id)
	if err != nil {
		println(err.Error())
	}
	return err
}
func (e *ElectionCommentDao) AddElectionComment(comment *model.ElectionComment) error {
	er := sqldb.Get(&comment.SubjectName, e.sql[6], comment.SubjectId)
	er = sqldb.Get(&comment.Classification, e.sql[9], comment.SubjectId)
	if er != nil {
		println(er.Error())
	}
	_, err := sqldb.Exec(e.sql[3], comment.SubjectId, comment.SubjectName, comment.Classification, comment.Comment)
	if err != nil {
		println(err.Error())
	}
	return err
}

func (e *ElectionCommentDao) GetMaxId() (maxID int, err error) {
	err = sqldb.Get(&maxID, "SELECT MAX(id) FROM election_comment")
	//println(maxID)
	return
}
