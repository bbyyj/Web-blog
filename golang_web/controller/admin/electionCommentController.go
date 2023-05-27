package admin

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

func (e *ElectionCommentController) ElectionCommentAllList(ctx *gin.Context) *response.Response {
	//subject_id := ctx.Query("subject_id")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	essays, err := e.electionCommentDao.FindAllElectionComment(pageNum, pageSize)
	if response.CheckError(err, "Get Election Comment List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}

func (e *ElectionCommentController) DeleteElectionComment(ctx *gin.Context) *response.Response {
	subject_id := ctx.Query("subject_id")
	err := e.electionCommentDao.DeleteElectionComment(subject_id)
	if response.CheckError(err, "Delete Election Comment error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
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
