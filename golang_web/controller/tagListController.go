package controller

import (
	"blog_web/db/service"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type TagListController struct {
	tagService  *service.TagService
	blogService *service.BlogService
}

func NewTagListRouter() *TagListController {
	return &TagListController{
		tagService:  service.NewTagService(),
		blogService: service.NewBlogService(),
	}
}

// 使用同后台一样的操作
func (t *TagListController) GetTagList(ctx *gin.Context) *response.Response {
	tags, err := t.tagService.GetAllTags()
	if response.CheckError(err, "Get tags error") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(tags)
}

// 使用blog里面的操作
func (t *TagListController) GetBlogListByTagId(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "8")
	id := utils.QueryInt(ctx, "tagId")

	blogs, i, err := t.blogService.GetBlogsByTagId(id, pageNum, pageSize)
	if response.CheckError(err, "get blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, i)
}
