package controller

import (
	"blog_web/db/service"
	"blog_web/response"
	"github.com/gin-gonic/gin"
)

type TacitController struct {
	tacitService *service.TacitService
}

func NewTacitRouter() *TacitController {
	return &TacitController{
		tacitService: service.NewTacitService(),
	}
}

func (e *TacitController) TacitList(ctx *gin.Context) *response.Response {
	essays, err := e.tacitService.GetAll()
	if response.CheckError(err, "Get Tacit List") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(essays)
}
