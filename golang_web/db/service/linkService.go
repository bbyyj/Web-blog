package service

import (
	"blog_web/db/dao"
	"blog_web/model"
)

type LinkService struct {
	linkDao *dao.LinkDao
}

func NewLinkService() *LinkService {
	return &LinkService{
		linkDao: dao.NewLinkDao(),
	}
}

func (l *LinkService) GetAllLinks() ([]model.Link, error) {
	return l.linkDao.FindAllLinks()
}

func (l *LinkService) GetAllCategory() ([]model.LinkCategory, error) {
	return l.linkDao.FindAllLinkCategory()
}

func (l *LinkService) GetLimitedLinks(pageNum, pageSize int) ([]model.Link, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.FindLimitedLinks(pageStart, pageSize)
}

func (l *LinkService) GetLinkCount() (int, error) {
	return l.linkDao.FindLinkCount()
}

func (l *LinkService) AddLink(link *model.Link) error {
	return l.linkDao.AddLink(link)
}

func (l *LinkService) DeleteLink(id int) error {
	return l.linkDao.DeleteLink(id)
}

func (l *LinkService) UpdateLink(link *model.Link) error {
	return l.linkDao.UpdateLink(link)
}

func (l *LinkService) AddCategory(category *model.LinkCategory) error {
	return l.linkDao.AddCategory(category)
}

func (l *LinkService) DeleteCategory(id int) error {
	return l.linkDao.DeleteCategory(id)
}

func (l *LinkService) UpdateCategory(category *model.LinkCategory) error {
	return l.linkDao.UpdateCategory(category)
}

func (l *LinkService) GetAllResource(pagestart int, pagesize int) ([]model.ResourceList, error) {
	return l.linkDao.FindAllResource(pagestart, pagesize)
}

func (l *LinkService) GetAllResourceByCategoryId(categoryId int, pagestart int, pagesize int) ([]model.ResourceManage, int, error) {
	return l.linkDao.FindAllResourceByCategoryId(categoryId, pagestart, pagesize)
}

func (l *LinkService) GetAllResourceLikeName(name string, pageNum, pageSize int) ([]model.ResourceList, int, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.FindAllResourceLikeName(name, pageStart, pageSize)
}

func (l *LinkService) GetDownloadUrl(name string) (*model.ResourceUrl, error) {
	return l.linkDao.GetDownloadUrl(name)
}

func (l *LinkService) GetAllResourceByName(name string) (model.ResourceManage, error) {
	return l.linkDao.FindAllResourceByName(name)
}

func (l *LinkService) UpdateResourceDownloadNumByName(name string) error {
	return l.linkDao.UpdateResourceDownloadNumByName(name)
}

func (l *LinkService) GetAllCategoryNew() ([]model.LinkCategory, error) {
	return l.linkDao.FindAllCategory()
}

//admin

func (l *LinkService) GetLimitedResource(pageNum, pageSize int) ([]model.ResourceManage, int, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.FindLimitedResource(pageStart, pageSize)
}

func (l *LinkService) GetLimitedResourceCheck(pageNum, pageSize int) ([]model.ResourceManage, int, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.FindLimitedResourceCheck(pageStart, pageSize)
}

func (l *LinkService) UpdateResource(link *model.ResourceManage) error {
	return l.linkDao.UpdateResource(link)
}

func (l *LinkService) CheckResourceName(id uint, name string) error {
	return l.linkDao.CheckResourceName(id, name)
}

func (l *LinkService) CheckResourceCheckName(id uint, name string) error {
	return l.linkDao.CheckResourceCheckName(id, name)
}

func (l *LinkService) AddResource(link *model.ResourceManage) error {
	return l.linkDao.AddResource(link)
}

func (l *LinkService) AddResourceCheck(link *model.ResourceManage) error {
	return l.linkDao.AddResourceCheck(link)
}

func (l *LinkService) DeleteResource(id uint) error {
	return l.linkDao.DeleteResource(id)
}

func (l *LinkService) GetResourceLikeName(name string, pageNum, pageSize int) ([]model.ResourceManage, int, error) {
	pageStart := (pageNum - 1) * pageSize
	return l.linkDao.GetResourceLikeName(name, pageStart, pageSize)
}

func (l *LinkService) ReUploadUpdateTime(name string) error {
	return l.linkDao.ReUploadUpdateTime(name)
}

func (l *LinkService) GetResourceUrl(id uint) *model.ResourceUrl {
	return l.linkDao.GetResourceUrl(id)
}

func (l *LinkService) GetCheckResourceById(id uint) *model.ResourceManage {
	return l.linkDao.GetCheckResourceById(id)
}

func (l *LinkService) DeleteResourceCheck(id uint) error {
	return l.linkDao.DeleteResourceCheck(id)
}
