package controller

import (
	"blog_web/db/dao"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type ElectionController struct {
	electionDao *dao.ElectionDao
}

func NewElectionRouter() *ElectionController {
	return &ElectionController{
		electionDao: dao.NewElectionDao(),
	}
}
func (e *ElectionController) ElectionList(ctx *gin.Context) *response.Response {
	classification := ctx.Query("classification")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	essays, count, err := e.electionDao.FindAllElection(classification, pageNum, pageSize)
	if response.CheckError(err, "Get Election List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays, count)
}
func (e *ElectionController) ElectionListNoClass(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	essays, count, err := e.electionDao.FindAllElectionNoClass(pageNum, pageSize)
	if response.CheckError(err, "Get Election List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays, count)
}

func (e *ElectionController) ElectionDetailedList(ctx *gin.Context) *response.Response {
	subjectId := ctx.Query("subject_id")
	essays, err := e.electionDao.FindDetailedElection(subjectId)
	if response.CheckError(err, "Get ElectionDetailed List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}
