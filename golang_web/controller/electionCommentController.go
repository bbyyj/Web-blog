package controller

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ElectionCommentController struct {
	electionCommentDao *dao.ElectionCommentDao
}

func NewElectionCommentRouter() *ElectionCommentController {
	return &ElectionCommentController{
		electionCommentDao: dao.NewElectionCommentDao(),
	}
}
func (e *ElectionCommentController) ElectionCommentList(ctx *gin.Context) *response.Response {
	subject_id := ctx.Query("subject_id")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pegeSize := utils.QueryInt(ctx, "pageSize")
	essays, count, err := e.electionCommentDao.FindElectionComment(subject_id, pageNum, pegeSize)
	if response.CheckError(err, "Get Election Comment List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays, count)
}
func (e *ElectionCommentController) AddElectionComment(ctx *gin.Context) *response.Response {
	var comment model.ElectionComment
	err := ctx.ShouldBind(&comment)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.electionCommentDao.AddElectionComment(&comment)
	if response.CheckError(err, "Add Election Comment error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
