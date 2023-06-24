package dao

import (
	"blog_web/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type LinkDao struct {
	sql []string
}

func NewLinkDao() *LinkDao {
	return &LinkDao{
		sql: []string{
			"SELECT id, `name`, `desc`, url, category_id, icon FROM t_resource_link;",
			"SELECT id, `name` FROM t_resource_category;",
			"SELECT id, `name`, `desc`, url, category_id, icon FROM t_resource_link LIMIT ?, ?;",
			"SELECT COUNT(*) FROM t_resource_link;",
			"INSERT INTO t_resource_link (`name`, `desc`, url, category_id, icon) VALUES (?, ?, ?, ?, ?);",
			"DELETE FROM t_resource_link WHERE id = ?",
			"UPDATE t_resource_link SET `name` = ?, `desc` = ?, url = ?, category_id = ?, icon = ? WHERE id = ?;",
			"INSERT INTO t_resource_category (`name`) VALUES (?);",
			"DELETE FROM t_resource_category WHERE id = ?;",
			"UPDATE t_resource_category SET `name` = ? WHERE id = ?;",
		},
	}
}

func (l *LinkDao) FindAllLinks() (links []model.Link, err error) {
	err = sqldb.Select(&links, l.sql[0])
	return
}

func (l *LinkDao) FindAllLinkCategory() (categories []model.LinkCategory, err error) {
	err = sqldb.Select(&categories, l.sql[1])
	return
}

func (l *LinkDao) FindLimitedLinks(pageStart, pageSize int) (links []model.Link, err error) {
	err = sqldb.Select(&links, l.sql[2], pageStart, pageSize)
	return
}

func (l *LinkDao) FindLinkCount() (count int, err error) {
	err = sqldb.Get(&count, l.sql[3])
	return
}

func (l *LinkDao) AddLink(link *model.Link) error {
	_, err := sqldb.Exec(l.sql[4], link.Name, link.Desc, link.Url, link.CategoryId, link.Icon)
	return err
}

func (l *LinkDao) DeleteLink(id int) error {
	_, err := sqldb.Exec(l.sql[5], id)
	return err
}

func (l *LinkDao) UpdateLink(link *model.Link) error {
	_, err := sqldb.Exec(l.sql[6], link.Name, link.Desc, link.Url, link.CategoryId, link.Icon, link.Id)
	return err
}

func (l *LinkDao) AddCategory(category *model.LinkCategory) error {
	_, err := sqldb.Exec(l.sql[7], category.Name)
	return err
}

func (l *LinkDao) DeleteCategory(id int) error {
	_, err := sqldb.Exec(l.sql[8], id)
	return err
}

func (l *LinkDao) UpdateCategory(category *model.LinkCategory) error {
	_, err := sqldb.Exec(l.sql[9], category.Name, category.Id)
	return err
}

func (l *LinkDao) FindAllResource(pagestart int, pagesize int) (links []model.ResourceList, err error) {
	err = Db.Table("resource_manages").Offset((pagestart - 1) * pagesize).Limit(pagesize).Select("`name`,`desc`,`category_id`,`download_num`,`file_size`,`created_at`").Find(&links).Error
	if len(links) == 0 {
		err = errors.New("No matching records found")
	}
	fmt.Println(err)
	fmt.Println(links)
	return
}

func (l *LinkDao) FindAllResourceByCategoryId(categoryId int, pagestart int, pagesize int) (links []model.ResourceManage, count int, err error) {
	err = Db.Table("resource_manages").Where("category_id = ?", categoryId).Offset((pagestart - 1) * pagesize).Limit(pagesize).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	fmt.Println(err)
	fmt.Println((pagestart - 1) * pagesize)
	fmt.Println(pagesize)
	fmt.Println(count)
	fmt.Println(links)
	if len(links) == 0 {
		err = errors.New("No matching records found")
	}
	return
}

func (l *LinkDao) FindAllResourceLikeName(name string, pagestart, pagesize int) (links []model.ResourceList, count int, err error) {
	fmt.Println(pagestart)
	fmt.Println(pagesize)
	err = Db.Table("resource_manages").Select("`name`,`desc`,`category_id`,`download_num`,`file_size`,`created_at`").Where("name LIKE ?", name).Count(&count).Offset(pagestart).Limit(pagesize).Find(&links).Error
	//err = Db.Table("links").Where("name LIKE ?", name).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	fmt.Println(links)
	if len(links) == 0 {
		err = errors.New("No matching records found")
	}
	return
}

func (l *LinkDao) GetDownloadUrl(name string) (*model.ResourceUrl, error) {
	var url model.ResourceUrl
	err := Db.Table("resource_manages").Select("name,url").Where("name = ?", name).Find(&url).Error
	return &url, err
}

func (l *LinkDao) FindAllResourceByName(name string) (links model.ResourceManage, err error) {
	err = Db.Table("resource_manages").Limit(1).Where("name = ?", name).First(&links).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("No matching records found")
	}
	fmt.Println(err)
	fmt.Println(links)
	return
}

func (l *LinkDao) UpdateResourceDownloadNumByName(name string) (err error) {

	err = Db.Table("resource_manages").Where("name = ?", name).UpdateColumn("download_num", gorm.Expr("download_num + ?", 1)).Error

	fmt.Println(err)
	return
}

func (l *LinkDao) FindAllCategory() (category []model.LinkCategory, err error) {
	err = Db.Table("t_resource_category").Find(&category).Error
	if len(category) == 0 {
		err = errors.New("No matching records found")
	}
	fmt.Println(err)
	fmt.Println(category)
	return
}

// admin
func (l *LinkDao) FindLimitedResource(pageStart, pageSize int) (links []model.ResourceManage, count int, err error) {
	//err = Db.Table("resource_manages").Offset(pageStart).Limit(pageSize).Select("`name`,`desc`,`category_id`,`download_num`,`file_size`,`created_at`").Find(&links).Error
	fmt.Println(pageStart)
	fmt.Println(pageSize)
	var link []model.ResourceManage
	Db.Table("resource_manages").Find(&link).Count(&count)
	if pageStart < count {
		err = Db.Table("resource_manages").Offset(pageStart).Limit(pageSize).Find(&links).Error
	}
	//err = Db.Table("resource_manages").Offset(pageStart).Limit(pageSize).Find(&links).Error

	if len(links) == 0 {
		err = errors.New("No matching records found")
	}
	fmt.Println(err)
	//fmt.Println(links)
	return
}

func (l *LinkDao) FindLimitedResourceCheck(pageStart, pageSize int) (links []model.ResourceManage, count int, err error) {
	//err = Db.Table("resource_manages").Offset(pageStart).Limit(pageSize).Select("`name`,`desc`,`category_id`,`download_num`,`file_size`,`created_at`").Find(&links).Error
	fmt.Println(pageStart)
	fmt.Println(pageSize)
	var link []model.ResourceManage
	Db.Table("resource_manage_checks").Find(&link).Count(&count)
	if pageStart < count {
		err = Db.Table("resource_manage_checks").Offset(pageStart).Limit(pageSize).Find(&links).Error
	}
	//err = Db.Table("resource_manages").Offset(pageStart).Limit(pageSize).Find(&links).Error

	//if len(links) == 0 {
	//	err = errors.New("No matching records found")
	//}
	fmt.Println(err)
	//fmt.Println(links)
	return
}

func (l *LinkDao) UpdateResource(link *model.ResourceManage) error {
	var err error
	fmt.Println(link.ID)
	result := Db.Model(&model.ResourceManage{}).Table("resource_manages").Where("id = ?", link.ID).Updates(&link)
	if result.RowsAffected == 0 {
		err = errors.New("No matching records found")
	}
	fmt.Println(err)
	return err
}

func (l *LinkDao) CheckResourceName(id uint, name string) error {
	var u []model.ResourceManage
	err := Db.Table("resource_manages").Where("name = ? AND id <> ?", name, id).Find(&u).Error
	//err = Db.Table("links").Where("name LIKE ?", name).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	//fmt.Println(u)
	if len(u) != 0 {
		return errors.New("Find same name!")
	}
	//fmt.Println(err)
	return err
}

func (l *LinkDao) CheckResourceCheckName(id uint, name string) error {
	var u []model.ResourceManage
	err := Db.Table("resource_manage_checks").Where("name = ? AND id <> ?", name, id).Find(&u).Error
	//err = Db.Table("links").Where("name LIKE ?", name).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	//fmt.Println(u)
	if len(u) != 0 {
		return errors.New("Find same name!")
	}
	//fmt.Println(err)
	return err
}

func (l *LinkDao) AddResource(resource *model.ResourceManage) error {
	var err error
	err = Db.Table("resource_manages").Create(&resource).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return err
}

func (l *LinkDao) AddResourceCheck(resource *model.ResourceManage) error {
	var err error
	err = Db.Table("resource_manage_checks").Create(&resource).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return err
}

func (l *LinkDao) DeleteResource(id uint) error {
	var err error
	err = Db.Table("resource_manages").Where("id = ?", id).Unscoped().Delete(model.ResourceManage{}).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return err
}

func (l *LinkDao) DeleteResourceCheck(id uint) error {
	var err error
	err = Db.Table("resource_manage_checks").Where("id = ?", id).Unscoped().Delete(model.ResourceManage{}).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return err
}

func (l *LinkDao) GetResourceLikeName(name string, pagestart, pagesize int) (resource []model.ResourceManage, count int, err error) {
	name = ("%" + name + "%")
	err = Db.Table("resource_manages").Where("name LIKE ?", name).Count(&count).Offset(pagestart).Limit(pagesize).Find(&resource).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return
}

func (l *LinkDao) ReUploadUpdateTime(name string) error {
	err := Db.Table("resource_manages").Where("name = ?", name).UpdateColumn("updated_at", time.Now()).Error
	//fmt.Println(u)
	//fmt.Println(err)
	return err
}

func (l *LinkDao) GetResourceUrl(id uint) *model.ResourceUrl {
	var url model.ResourceUrl
	Db.Table("resource_manages").Select("name,url").Where("id = ?", id).Find(&url)
	//err = Db.Table("links").Where("name LIKE ?", name).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	//fmt.Println(url)
	//fmt.Println(err)
	//fmt.Println(count)
	return &url
}

func (l *LinkDao) GetCheckResourceById(id uint) *model.ResourceManage {
	var u model.ResourceManage
	Db.Table("resource_manage_checks").Where("id = ?", id).Find(&u)
	//err = Db.Table("links").Where("name LIKE ?", name).Find(&links).Offset(-1).Limit(-1).Count(&count).Error
	//fmt.Println(url)
	//fmt.Println(err)
	//fmt.Println(count)
	return &u
}
