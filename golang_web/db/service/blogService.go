package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type BlogService struct {
	blogDao *dao.BlogDao
	tagDao  *dao.TagDao
}

func NewBlogService() *BlogService {
	return &BlogService{
		blogDao: dao.NewBlogDao(),
		tagDao:  dao.NewTagDao(),
	}
}

func (b *BlogService) GetHomePageBlogs(pageNum, pageSize int) ([]model.BlogUserType, error) {
	pageStart := (pageNum - 1) * pageSize
	return b.blogDao.GetConciseBlogs(pageStart, pageSize)
}

func (b *BlogService) GetBlogCount() (int, error) {
	return b.blogDao.GetBlogCount()
}

func (b *BlogService) GetDetailedBlog(id int) (*model.DetailedBlog, []model.Tag, error) {
	b.blogDao.UpdateViews(id)
	blog, err := b.blogDao.FindDetailedBlog(id)
	if err != nil {
		return nil, nil, err
	}
	tags, _ := b.tagDao.GetTagsByBlogId(id)

	return blog, tags, nil
}

func (b *BlogService) GetAllTypes() ([]int, error) {
	return b.blogDao.FindAllTypes()
}

func (b *BlogService) GetBlogListByTypeId(typeid, pageNum, pageSize int) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindByTypeId(typeid, pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}

	count, err := b.blogDao.GetBolgCountByTypeId(typeid)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

func (b *BlogService) GetBlogsByTagId(id, pageNum, pageSize int) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, err := b.blogDao.FindBlogsByTagId(id, pageStart, pageSize)
	if err != nil {
		return nil, 0, err
	}

	count, err := b.blogDao.GetBlogCountByTagId(id)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

func (b *BlogService) GetBlogsByKeyWord(keyWord string) ([]model.BlogSection, error) {
	keyWord = "%" + keyWord + "%"
	blogs, err := b.blogDao.FindBlogsByKeyWord(keyWord)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *BlogService) GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize int, title string, typeId int,
	recommend string) ([]model.BlogUserType, int, error) {
	pageStart := (pageNum - 1) * pageSize
	blogs, count, err := b.blogDao.FindBlogsByTitleOrTypeOrRecommend(pageStart, pageSize, title, typeId, recommend)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

func (b *BlogService) DeleteBlog(id int) error {
	return b.blogDao.DeleteById(id)
}

func (b *BlogService) GetFullBlog(id int) (*model.BlogUserType, []model.Tag, error) {
	blog, err := b.blogDao.FindFullBlog(id)
	if err != nil {
		return nil, nil, err
	}

	tags, err := b.tagDao.GetTagsByBlogId(id)
	if err != nil {
		return nil, nil, err
	}

	return blog, tags, nil
}

func (b *BlogService) UpdateBlog(blog *model.FullBlog) error {
	// 和addBlog事务操作类似
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		}
	}()

	//  blog_tags表 删除原本数据
	err = b.blogDao.DeleteBlogTagsByBlogId(tx, blog.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// blog表 修改
	err = b.blogDao.UpdateBlog(tx, blog)
	if err != nil {
		tx.Rollback()
		return err
	}

	// blog_tags表 插入
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = blog.Id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (b *BlogService) AddBlog(blog *model.FullBlog) error {

	// 多个表操作 开启一个数据库事务 出错则回滚全部撤销
	tx, err := dao.Sqldb.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			tx.Rollback()
		}
	}()

	// blog表
	id, err := b.blogDao.AddBlog(tx, blog)
	if err != nil {
		tx.Rollback()
		return err
	}
	// blog_tags表
	blogTags := make([]model.BlogTag, len(blog.TagIds))
	for i, tid := range blog.TagIds {
		blogTags[i].TagId = tid
		blogTags[i].BlogId = id
	}
	err = b.blogDao.InsertIntoBlogTags(tx, blogTags)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 提交一个数据库事务
	tx.Commit()

	return nil
}

func (b *BlogService) GetTypeAndBlogCount() ([]model.TheType, error) {
	return b.blogDao.FindTypeAndBlogCount()
}
