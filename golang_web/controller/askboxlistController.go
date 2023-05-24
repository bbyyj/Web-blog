package controller

import (
	"blog_web/db/service"
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
