package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService *service.TagService
}

func NewTagRouter() *TagController {
	return &TagController{
		tagService: service.NewTagService(),
	}
}

func (t *TagController) GetAllTags(ctx *gin.Context) *response.Response {
	tags, err := t.tagService.GetAllTags()
	if response.CheckError(err, "Get tags error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(tags)
}

func (t *TagController) GetOnePageTags(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	tags, count, err := t.tagService.GetOnePageTags(pageNum, pageSize)
	if response.CheckError(err, "Get one page tag error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(tags, count)
}

func (t *TagController) AddTag(ctx *gin.Context) *response.Response {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		return response.ROperateFailed()
	}

	err := t.tagService.AddTag(tag.Name)
	if response.CheckError(err, "Add tag error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (t *TagController) UpdateTag(ctx *gin.Context) *response.Response {
	var tag model.Tag
	if err := ctx.ShouldBind(&tag); err != nil {
		return response.ROperateFailed()
	}

	err := t.tagService.UpdateTagName(tag.Id, tag.Name)
	if response.CheckError(err, "Update tag error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (t *TagController) CheckTagExist(ctx *gin.Context) *response.Response {
	name := ctx.Query("name")
	exist := t.tagService.CheckTagExist(name)
	return response.RQuerySuccess(exist)
}

func (t *TagController) DeleteTag(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := t.tagService.DeleteTagById(id)
	if response.CheckError(err, "Delete tag error") {
		return response.RDeleteFailed()
	}

	return response.RDeleteSuccess()
}
