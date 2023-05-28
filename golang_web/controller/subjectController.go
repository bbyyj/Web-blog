package controller

import (
	"blog_web/db/dao"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	subjectDao *dao.SubjectDao
}

type ChapterController struct {
	chapterDao *dao.ChapterDao
}

type ExamController struct {
	examDao *dao.ExamDao
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
func NewExamRouter() *ExamController {
	return &ExamController{
		examDao: dao.NewExamDao(),
	}
}

func (e *SubjectController) SubjectList(ctx *gin.Context) *response.Response {
	essays, err := e.subjectDao.FindAll()
	if response.CheckError(err, "Get Subject List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(essays)
}

//func (e *SubjectController) UpdateSubject(ctx *gin.Context) *response.Response {
//	updatesub := struct {
//		Old_sub string `json:"old_subject"`
//		New_sub string `json:"new_subject"`
//	}{}
//	err := ctx.ShouldBind(&updatesub)
//	if response.CheckError(err, "Bind param error") {
//		ctx.Status(http.StatusInternalServerError)
//		return nil
//	}
//	err = e.subjectDao.UpdateSub(updatesub.Old_sub, updatesub.New_sub)
//	if response.CheckError(err, "Update subject error") {
//		return response.ResponseOperateFailed()
//	}
//	return response.ResponseOperateSuccess()
//
//}

func (e *ChapterController) ChapterList(ctx *gin.Context) *response.Response {
	name := ctx.Query("name")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	chapters, err := e.chapterDao.FindChapter(name, pageNum, pageSize)
	if response.CheckError(err, "Get Chapter List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(chapters)
}

func (e *ExamController) ExamList(ctx *gin.Context) *response.Response {
	a := ctx.Query("typename")
	b := ctx.Query("title")
	exams, err := e.examDao.FindExamLimited(a, b)
	if response.CheckError(err, "Get Exam List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(exams)

}
