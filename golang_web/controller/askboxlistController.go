package controller

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AskBoxFrontController struct {
	askBoxService *service.AskBoxService
}

func NewAskboxFrontRouter() *AskBoxFrontController {
	return &AskBoxFrontController{
		askBoxService: service.NewAskBoxService(),
	}
}

// 报错？
//func (a *AskBoxFrontController) GetAnsweredQA(ctx *gin.Context) *response.Response {
//	askboxs, err := a.askBoxService.GetAnsweredQA()
//	if response.CheckError(err, "Get AnsweredQA error") {
//		return response.ResponseQueryFailed()
//	}
//	return response.ResponseQuerySuccess(askboxs)
//}

func (a *AskBoxFrontController) AddNewQuestion(ctx *gin.Context) *response.Response {

	var askbox model.Askbox
	err := ctx.ShouldBind(&askbox)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	//`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
	// 已传入 parent_id, question, rainbow
	askbox.ChildId = 1
	askbox.QuestionTime = time.Now()
	askbox.IsParent = true
	err = a.askBoxService.AddNewQuestion(&askbox)

	if response.CheckError(err, "Add NewQuestion error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}

func (a *AskBoxFrontController) AppendOldQuestion(ctx *gin.Context) *response.Response {

	var askbox model.Askbox
	err := ctx.ShouldBind(&askbox)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	//`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow)
	// 已传入 parent_id, child_id, question, rainbow
	askbox.ChildId += 1
	askbox.QuestionTime = time.Now()
	err = a.askBoxService.AddNewQuestion(&askbox)

	if response.CheckError(err, "Append OldQuestion error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}

func (a *AskBoxFrontController) ClickLikes(ctx *gin.Context) *response.Response {
	var askbox model.Askbox
	if err := ctx.ShouldBind(&askbox); err != nil {
		return response.ResponseOperateFailed()
	}

	askbox.Likes += 1
	err := a.askBoxService.ClickLikes(askbox.Likes, askbox.ParentId, askbox.ChildId)

	if response.CheckError(err, "Click Likes error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
