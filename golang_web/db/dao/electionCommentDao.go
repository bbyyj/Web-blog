package dao

import "blog_web/model"

type ElectionCommentDao struct {
	sql []string
}

func NewElectionCommentDao() *ElectionCommentDao {
	return &ElectionCommentDao{
		sql: []string{
			`SELECT * FROM t_election_comment WHERE subject_id=? ORDER BY id ASC LIMIT ?,?;`,
			`SELECT * FROM t_election_comment ORDER BY id ASC LIMIT ?,?;`,
			`DELETE FROM t_election_comment WHERE id=?`,
			`INSERT INTO t_election_comment(subject_id,subject_name,comment) VALUES (?,?,?)`,
		},
	}
}
func (e *ElectionCommentDao) FindElectionComment(subjectId string, pageNum int, pageSize int) (electionCommentList []model.ElectionComment, err error) {
	err = sqldb.Select(&electionCommentList, e.sql[0], subjectId, pageNum, pageSize)
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
func (e *ElectionCommentDao) DeleteElectionComment(id int) error {
	_, err := sqldb.Exec(e.sql[2], id)
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
