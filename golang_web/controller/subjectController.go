package controller

import (
	"blog_web/db/dao"
	"blog_web/response"
	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	subjectDao *dao.SubjectDao
}

type ChapterController struct {
	chapterDao *dao.ChapterDao
}

func NewSubjectRouter() *SubjectController {
	return &SubjectController{
		subjectDao: dao.NewSubjectDao(),
	}
}

func NewChapterRouter() *ChapterController {
	return &ChapterController{
		chapterDao: dao.NewChapterDao(),
	}
}

func (e *SubjectController) SubjectList(ctx *gin.Context) *response.Response {
	essays, err := e.subjectDao.FindAll()
	if response.CheckError(err, "Get Subject List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}

func (e *ChapterController) ChapterList(ctx *gin.Context) *response.Response {
	s := ctx.Query("subject")
	essays, err := e.chapterDao.FindChapter(s)
	if response.CheckError(err, "Get Chapter List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}
