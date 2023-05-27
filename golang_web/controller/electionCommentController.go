package controller

import (
	"blog_web/db/dao"
	"blog_web/response"
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
func (e *ElectionCommentController) ElectionCommentList(ctx *gin.Context) *response.Response {
	subject_id := ctx.Query("subject_id")
	//pageNum := utils.QueryInt(ctx, "pageNum")
	//pegeSize := utils.QueryInt(ctx, "pageSize")
	essays, err := e.electionCommentDao.FindElectionComment(subject_id)
	if response.CheckError(err, "Get Election Comment List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}
