package dao

import (
	"blog_web/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type BlogDao struct {
	sql []string
}

func NewBlogDao() *BlogDao {
	return &BlogDao{
		sql: []string{
			// 分页返回首页精简文章
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
			b.views, u.nickname nickname, t.name typename
			FROM blog b, user u, type t
			WHERE b.user_id = u.id and b.type_id = t.id and b.published = 1
			ORDER BY b.create_time DESC
			LIMIT ?, ?`,
			//  返回全部文章数量
			"SELECT COUNT(*) FROM blog;",
			// 更新浏览次数
			"UPDATE blog SET views = views + 1 WHERE id = ?;",
			// 返回一篇文章详细内容
			`SELECT b.id, b.title, b.create_time, b.update_time, b.flag, b.appreciation, b.views, b.content, u.nickname nickname, t.name typename
			FROM blog b JOIN user u ON b.user_id = u.id 
			JOIN type t ON b.type_id = t.id 
			WHERE b.id = ?;`,
			// 返回所有文章已用的分类
			`SELECT type_id FROM blog;`,
			// 根据类型分页返回文章
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
			b.views, u.nickname nickname, t.name typename
        	FROM blog b, user u, type t
        	WHERE b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.type_id = ?
			ORDER BY b.create_time DESC
			LIMIT ?, ?`,
			// 返回一个类型下的文章数量
			`SELECT COUNT(*) FROM blog b, user u, type t WHERE 
			b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.type_id = ?;`,
			// 根据标签返回文章
			`SELECT b.id, b.title, b.description, b.first_picture, b.create_time,
        	b.views, u.nickname nickname, t.name typename
        	FROM blog b, user u, type t
        	WHERE b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.id in (
            SELECT blog_id from blog_tags WHERE tag_id = ?)
			ORDER BY b.create_time DESC 
			LIMIT ?, ?;`,
			// 返回一个标签下的文章数量
			`SELECT COUNT(*) FROM blog b, user u, type t WHERE
			b.user_id = u.id AND b.type_id = t.id AND b.published = 1 AND b.id in (
            SELECT blog_id FROM blog_tags WHERE tag_id = ?)`,
			// 9、查询时间线
			`SELECT DISTINCT DATE_FORMAT(blog.create_time, '%Y') ut FROM blog ORDER BY ut DESC;`,
			// 10、根据某年查询该年发布的博客
			`SELECT b.title, b.create_time, b.id, b.flag
        	FROM blog b WHERE DATE_FORMAT(b.create_time, "%Y") = ? ORDER BY b.create_time DESC;`,
			// 根据标题或内容搜索
			`SELECT id, title
			FROM blog WHERE title LIKE ? OR content LIKE ?
			ORDER BY create_time DESC;`,
			// 返回搜索的文章的数量
			`SELECT COUNT(*)
			FROM blog
			WHERE title LIKE ? OR content LIKE ?;`,
			// 根据标题、类型、推荐返回详细文章和文章类型（需要拼接sql语句）
			`SELECT b.id, b.title, b.create_time, b.recommend, b.published, b.type_id, t.name typename 
			FROM blog b ,type t WHERE  b.type_id = t.id`,
			// 根据标题、类型、推荐返回文章数量
			`SELECT COUNT(*) FROM blog b, type t WHERE  b.type_id = t.id `,
			// 删除一篇文章
			`DELETE FROM blog WHERE id = ?;`,
			// 查询一篇详细文章
			`SELECT b.id, b.first_picture, b.flag, b.title, b.content, b.views,
        	b.update_time, b.commentabled, b.recommend, b.share_statement, b.appreciation, b.description,
        	t.name typename, t.id type_id
        	FROM blog b, type t
        	WHERE b.type_id = t.id AND  b.id = ?;`,
			// 更新一篇文章
			`UPDATE blog SET title=?, content=?, first_picture=?, flag=?, appreciation=?, share_statement=?,
			commentabled=?, published=?, recommend=?, update_time=?, type_id=?, user_id=?, description=? WHERE id=?`,
			// 新增一篇文章
			`INSERT INTO blog(title, content, first_picture, flag, views,
        	appreciation, share_statement, commentabled, published, recommend,
        	create_time, update_time, type_id, user_id, description) VALUES(?, ?,
        	?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?,?, ?)`,
			// 添加文章标签
			`INSERT INTO blog_tags (blog_id, tag_id) VALUES (:blog_id, :tag_id)`,
			// 删除文章对应标签
			`DELETE FROM blog_tags WHERE blog_id = ?;`,
			// 根据博客类型返回分组文章及分组内数量
			`SELECT t.id, count(*) 'count',  t.name FROM blog b JOIN type t ON b.type_id = t.id GROUP BY type_id;`,
		},
	}
}

func (b *BlogDao) GetConciseBlogs(pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[0], pageStart, pageSize)
	return
}

func (b *BlogDao) GetBlogCount() (count int, err error) {
	err = sqldb.Get(&count, b.sql[1])
	return
}

func (b *BlogDao) UpdateViews(id int) error {
	_, err := sqldb.Exec(b.sql[2], id)
	return err
}

func (b *BlogDao) FindDetailedBlog(id int) (*model.DetailedBlog, error) {
	var blog model.DetailedBlog
	err := sqldb.Get(&blog, b.sql[3], id)
	return &blog, err
}

func (b *BlogDao) FindAllTypes() (typeIds []int, err error) {
	err = sqldb.Select(&typeIds, b.sql[4])
	return
}

func (b *BlogDao) FindByTypeId(id, pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[5], id, pageStart, pageSize)
	return
}

func (b *BlogDao) GetBolgCountByTypeId(id int) (count int, err error) {
	err = sqldb.Get(&count, b.sql[6], id)
	return
}

func (b *BlogDao) FindBlogsByTagId(id, pageStart, pageSize int) (blogs []model.BlogUserType, err error) {
	err = sqldb.Select(&blogs, b.sql[7], id, pageStart, pageSize)
	return
}

func (b *BlogDao) GetBlogCountByTagId(id int) (count int, err error) {
	err = sqldb.Get(&count, b.sql[8], id)
	return
}

// 9、获取去重后的博客发布时间
func (b *BlogDao) GetAllBlogPublishYear() (years []string, err error) {
	err = sqldb.Select(&years, b.sql[9])
	return
}

// 10、根据%Y类型的时间获取博客
func (b *BlogDao) GetBlogByFormatedYear(year string) (blogs []model.Blog, err error) {
	err = sqldb.Select(&blogs, b.sql[10], year)
	return
}

// 11、根据关键词查询博客
func (b *BlogDao) FindBlogsByKeyWord(keyWord string) (blogs []model.BlogSection, err error) {
	err = sqldb.Select(&blogs, b.sql[11], keyWord, keyWord)
	return
}

// 12、根据关键词查询博客数量
func (b *BlogDao) GetBlogCountByKeyWord(keyWord string) (count int, err error) {
	err = sqldb.Get(&count, b.sql[12], keyWord, keyWord)
	return
}

// 13、根据三个关键字的组合来获取博客，分别为博客标题、博客类型以及是否推荐
func (b *BlogDao) FindBlogsByTitleOrTypeOrRecommend(pageStart, pageSize int, title string, typeId int, recommend string) (blogs []model.BlogUserType, count int, err error) {
	sq1 := b.sql[13]
	sq2 := b.sql[14]
	param := make([]interface{}, 0, 5)
	// `SELECT b.id, b.title, b.create_time, b.recommend, b.published, b.type_id, t.name typename
	//		FROM blog b ,type t WHERE  b.type_id = t.id`, 尾部拼接
	// fmt.Sprintf函数根据提供的格式字符串以及参数来生成一个新的字符串
	if title != "" {
		title = "%" + title + "%"
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.title like ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.title like ?")
		param = append(param, title)
	}
	if typeId > 0 {
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.type_id = ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.type_id = ?")
		param = append(param, typeId)
	}
	if recommend != "" {
		sq1 = fmt.Sprintf("%s %s", sq1, "AND b.recommend = ?")
		sq2 = fmt.Sprintf("%s %s", sq2, "AND b.recommend = ?")
		rec, err := strconv.ParseBool(recommend)
		if err != nil {
			return nil, 0, err
		}
		param = append(param, rec)
	}

	// 调整顺序
	sq1 = fmt.Sprintf("%s %s;", sq1, "ORDER BY b.create_time DESC LIMIT ?, ?")
	param = append(param, pageStart)
	param = append(param, pageSize)

	// sqldb.Select和sqldb.Get 接收多个参数而不是一个切片, 所以使用param...和par...来展开切片
	err = sqldb.Select(&blogs, sq1, param...)
	if err != nil {
		return nil, 0, err
	}

	par := param[:len(param)-2]
	err = sqldb.Get(&count, sq2, par...)
	if err != nil {
		return nil, 0, err
	}

	return
}

// 15、根据id删除博客
func (b *BlogDao) DeleteById(id int) error {
	_, err := sqldb.Exec(b.sql[15], id)
	return err
}

// 16、博客编辑页获取博客完整信息
func (b *BlogDao) FindFullBlog(id int) (*model.BlogUserType, error) {
	var blog model.BlogUserType
	err := sqldb.Get(&blog, b.sql[16], id)
	return &blog, err
}

// 17、更新博客
func (b *BlogDao) UpdateBlog(tx *sqlx.Tx, blog *model.FullBlog) error {
	_, err := tx.Exec(b.sql[17], blog.Title, blog.Content, blog.FirstPicture, blog.Flag, blog.Appreciation, blog.ShareStatement,
		blog.Commentabled, blog.Published, blog.Recommend, blog.UpdateTime, blog.TypeId, blog.UserId, blog.Description, blog.Id)
	return err
}

// 18、新增博客
func (b *BlogDao) AddBlog(tx *sqlx.Tx, blog *model.FullBlog) (int, error) {
	result, err := tx.Exec(b.sql[18], blog.Title, blog.Content, blog.FirstPicture, blog.Flag, blog.Views,
		blog.Appreciation, blog.ShareStatement, blog.Commentabled, blog.Published, blog.Recommend,
		blog.CreateTime, blog.UpdateTime, blog.TypeId, blog.UserId, blog.Description)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// 19、添加博客标签
func (b *BlogDao) InsertIntoBlogTags(tx *sqlx.Tx, bt []model.BlogTag) error {
	_, err := tx.NamedExec(b.sql[19], bt)
	return err
}

// 20、删除博客标签
func (b *BlogDao) DeleteBlogTagsByBlogId(tx *sqlx.Tx, id int) error {
	_, err := tx.Exec(b.sql[20], id)
	return err
}

// 21、根据博客类型对博客分组、计算每个分组的数量
func (b *BlogDao) FindTypeAndBlogCount() (types []model.TheType, err error) {
	err = sqldb.Select(&types, b.sql[21])
	return
}
