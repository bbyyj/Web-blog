package dao

import "blog_web/model"

type ElectionCommentDao struct {
	sql []string
}

func NewElectionCommentDao() *ElectionCommentDao {
	return &ElectionCommentDao{
		sql: []string{
			`SELECT * FROM t_election_comment WHERE subject_id=? ORDER BY id ASC;`,
			`SELECT * FROM t_election_comment ORDER BY id ASC LIMIT ?,?;`,
			`DELETE FROM t_election_comment WHERE subject_id=?`,
			`INSERT INTO t_election_comment(subject_id,subject_name,comment) VALUES (?,?,?)`,
		},
	}
}
func (e *ElectionCommentDao) FindElectionComment(subjectId string) (electionCommentList []model.ElectionComment, err error) {
	err = sqldb.Select(&electionCommentList, e.sql[0], subjectId)
	return
}

func (e *ElectionCommentDao) FindAllElectionComment(pageNum int, pageSize int) (electionCommentList []model.ElectionComment, err error) {
	pageStart := (pageNum - 1) * pageSize
	err = sqldb.Select(&electionCommentList, e.sql[1], pageStart, pageSize)
	if err != nil {
		println(err.Error())
	}
	return
}
func (e *ElectionCommentDao) DeleteElectionComment(subjectId string) error {
	_, err := sqldb.Exec(e.sql[2], subjectId)
	if err != nil {
		println(err.Error())
	}
	return err
}
func (e *ElectionCommentDao) AddElectionComment(comment *model.ElectionComment) error {
	_, err := sqldb.Exec(e.sql[3], comment.SubjectId, comment.SubjectName, comment.Comment)
	if err != nil {
		println(err.Error())
	}
	return err
}
