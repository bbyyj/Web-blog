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

func (a *AskBoxFrontController) GetAnsweredQA(ctx *gin.Context) *response.Response {
	askboxs, count, err := a.askBoxService.GetAnsweredQA()
	if response.CheckError(err, "Get AnsweredQA error") {
		return response.RQueryFailed()
	}
	return response.RQuerySuccess(askboxs, count)
}

func (a *AskBoxFrontController) AddNewQuestion(ctx *gin.Context) *response.Response {

	var askbox model.Askbox
	err := ctx.ShouldBind(&askbox)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	maxParentQuestionID, err := a.askBoxService.GetMaxParentQuestionId()
	if response.CheckError(err, "Get MaxParentQuestionId error") {
		return response.RQueryFailed()
	}
	askbox.ParentId = maxParentQuestionID + 1

	//`INSERT INTO t_askbox(parent_id, child_id, question, question_time, rainbow, is_parent)
	// 已传入 question, rainbow
	askbox.ChildId = 1
	askbox.QuestionTime = time.Now()
	askbox.IsParent = true
	err = a.askBoxService.AddNewQuestion(&askbox)

	if response.CheckError(err, "Add NewQuestion error") {
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
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
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}

func (a *AskBoxFrontController) ClickLikes(ctx *gin.Context) *response.Response {
	var askbox model.Askbox
	if err := ctx.ShouldBind(&askbox); err != nil {
		return response.ROperateFailed()
	}
	println(askbox.Likes)
	//askbox.Likes += 1
	err := a.askBoxService.ClickLikes(askbox.Likes, askbox.ParentId, askbox.ChildId)

	if response.CheckError(err, "Click Likes error") {
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}
