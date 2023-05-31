package admin

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ElectionController struct {
	electionDao *dao.ElectionDao
}

func NewElectionRouter() *ElectionController {
	return &ElectionController{
		electionDao: dao.NewElectionDao(),
	}
}
func (e *ElectionController) ElectionListNoPage(ctx *gin.Context) *response.Response {
	classification := ctx.Query("classification")
	essays, err := e.electionDao.FindAllElectionNopage(classification)
	if response.CheckError(err, "Get Election List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}

func (e *ElectionController) ElectionByClass(ctx *gin.Context) *response.Response {
	classification := ctx.Query("classification")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	essays, count, err := e.electionDao.FindElectionByClass(classification, pageNum, pageSize)
	if response.CheckError(err, "Get Election List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays, count)
}
func (e *ElectionController) ElectionAllList(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	essays, count, err := e.electionDao.FindAllDetailedElection(pageNum, pageSize)
	if response.CheckError(err, "Get Election List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays, count)
}
func (e *ElectionController) DeleteElection(ctx *gin.Context) *response.Response {
	subjectId := ctx.Query("subject_id")
	err := e.electionDao.DeleteElection(subjectId)
	if response.CheckError(err, "Delete exam error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseDeleteSuccess()
}
func (e *ElectionController) AddElection(ctx *gin.Context) *response.Response {
	var election model.ElectionDetailed
	err := ctx.ShouldBind(&election)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.electionDao.CreateElection(&election)
	if response.CheckError(err, "Add Election error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
func (e *ElectionController) UpdateElection(ctx *gin.Context) *response.Response {
	var election model.ElectionDetailed
	err := ctx.ShouldBind(&election)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.electionDao.UpdateElection(&election)
	if response.CheckError(err, "Add Election error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
