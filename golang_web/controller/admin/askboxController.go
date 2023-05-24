package admin

import (
	"blog_web/db/service"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type AskBoxBackController struct {
	askBoxService *service.AskBoxService
}

func NewAskBoxBackRouter() *AskBoxBackController {
	return &AskBoxBackController{
		askBoxService: service.NewAskBoxService(),
	}
}

func (a *AskBoxBackController) GetAllQA(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	messages, err := a.askBoxService.GetAllQA(pageNum, pageSize)

	if response.CheckError(err, "Get messages error") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(messages)
}

// 分页返回所有未回答的问题信息
func (a *AskBoxBackController) GetUnansweredQA(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	messages, err := a.askBoxService.GetUnansweredQA(pageNum, pageSize)

	if response.CheckError(err, "Find Unanswered QA error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(messages)
}

//func (a *AskBoxBackController) AddAnswer(ctx *gin.Context) *response.Response {
//	var msg model.Askbox
//	err := ctx.ShouldBind(&msg)
//	if response.CheckError(err, "Bind param error") {
//		ctx.Status(http.StatusInternalServerError)
//		return nil
//	}
//
//	msg.IsAnswered = true
//	msg.AnswerTime = time.Now()
//
//	err = a.askBoxService.AddAnswer(&msg)
//	if response.CheckError(err, "Add answer error") {
//		return response.ResponseOperateFailed()
//	}
//
//	return response.ResponseOperateSuccess()
//}
//
//func (a *AskBoxBackController) ModifyAnswer(ctx *gin.Context) *response.Response {
//	var msg model.Askbox
//	err := ctx.ShouldBind(&msg)
//	if response.CheckError(err, "Bind param error") {
//		ctx.Status(http.StatusInternalServerError)
//		return nil
//	}
//
//	msg.IsAnswered = true
//	msg.AnswerTime = time.Now()
//
//	err = a.askBoxService.ModifyAnswer(&msg)
//	if response.CheckError(err, "Modify answer error") {
//		return response.ResponseOperateFailed()
//	}
//
//	return response.ResponseOperateSuccess()
//}
//
//func (a *AskBoxBackController) DeleteQuestion(ctx *gin.Context) *response.Response {
//	id := utils.QueryInt(ctx, "parent_id")
//	err := a.askBoxService.DeleteQuestion(id)
//	if response.CheckError(err, "Delete Question error") {
//		return response.ResponseDeleteFailed()
//	}
//
//	return response.ResponseDeleteSuccess()
//}
//
