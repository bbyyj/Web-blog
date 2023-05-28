package admin

import (
	"blog_web/db/dao"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
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
	id := utils.QueryInt(ctx, "id")
	err := e.electionCommentDao.DeleteElectionComment(id)
	if response.CheckError(err, "Delete Election Comment error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
