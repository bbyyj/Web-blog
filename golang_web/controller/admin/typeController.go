package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type TypeController struct {
	typeService *service.TypeService
}

func NewTypeRouter() *TypeController {
	return &TypeController{
		typeService: service.NewTypeService(),
	}
}

func (t *TypeController) GetAllTypes(ctx *gin.Context) *response.Response {
	types, err := t.typeService.FindAll()
	if response.CheckError(err, "Find types error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(types)
}

func (t *TypeController) GetOnePageTypes(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	types, count, err := t.typeService.GetOnePage(pageNum, pageSize)
	if response.CheckError(err, "Get one page types error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(types, count)
}

func (t *TypeController) CheckTypeExist(ctx *gin.Context) *response.Response {
	typename := ctx.Query("typename")
	exist := t.typeService.CheckTypeExist(typename)
	return response.RQuerySuccess(exist)
}

func (t *TypeController) DeleteType(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := t.typeService.DeleteById(id)
	if response.CheckError(err, "Delete type error") {
		return response.RDeleteFailed()
	}

	return response.RDeleteSuccess()
}

func (t *TypeController) UpdateType(ctx *gin.Context) *response.Response {
	var tp model.TheType
	err := ctx.ShouldBind(&tp)
	if response.CheckError(err, "Bind param error") {
		return response.ROperateFailed()
	}

	err = t.typeService.UpdateName(tp.Id, tp.Name)
	if response.CheckError(err, "Update type error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (t *TypeController) AddType(ctx *gin.Context) *response.Response {
	var tp model.TheType
	ctx.ShouldBind(&tp)
	err := t.typeService.AddType(tp.Name)
	if response.CheckError(err, "Add type error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}
