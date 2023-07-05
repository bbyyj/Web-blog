package dao

import "blog_web/model"

type TagDao struct {
	sql []string
}

func NewTagDao() *TagDao {
	return &TagDao{
		sql: []string{
			// FROM blog_tags bt JOIN tag t ON t.id = bt.tag_id
			// blog_tags表与tag表通过它们之间的tag_id和id字段进行连接 （一个博客有多个标签）
			`SELECT t.id, t.name FROM blog_tags bt JOIN tag t ON t.id = bt.tag_id WHERE bt.blog_id = ?;`,
			`SELECT * FROM tag ;`,
			`SELECT * FROM tag LIMIT ?, ?;`,
			`SELECT * FROM tag WHERE name = ?;`,
			`DELETE FROM tag WHERE id = ?;`,
			`UPDATE tag SET name = ? WHERE id = ?;`,
			`INSERT INTO tag (name) VALUES(?);`,
			`SELECT COUNT(*) FROM tag;`,
		},
	}
}

func (t *TagDao) GetTagsByBlogId(id int) (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[0], id)
	return
}

func (t *TagDao) FindAllTags() (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[1])
	return
}

func (t *TagDao) GetOnePageTags(pageStart, pageSize int) (tags []model.Tag, err error) {
	err = sqldb.Select(&tags, t.sql[2], pageStart, pageSize)
	return
}

func (t *TagDao) FindByName(name string) (*model.Tag, error) {
	var tag model.Tag
	err := sqldb.Get(&tag, t.sql[3], name)
	return &tag, err
}

func (t *TagDao) DeleteById(id int) error {
	_, err := sqldb.Exec(t.sql[4], id)
	return err
}

func (t *TagDao) UpdateName(id int, name string) error {
	_, err := sqldb.Exec(t.sql[5], name, id)
	return err
}

func (t *TagDao) AddTag(name string) error {
	_, err := sqldb.Exec(t.sql[6], name)
	return err
}

func (t *TagDao) GetTagsCount() (int, error) {
	var count int
	err := sqldb.Get(&count, t.sql[7])
	return count, err
}
