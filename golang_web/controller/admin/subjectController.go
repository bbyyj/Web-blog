package admin

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//学科部分的后台管理
type SubjectController struct {
	subjectDao *dao.SubjectDao
}

func NewSubjectRouter() *SubjectController {
	return &SubjectController{
		subjectDao: dao.NewSubjectDao(),
	}
}

func (e *SubjectController) AddSubject(ctx *gin.Context) *response.Response {
	var subject model.Subject
	err := ctx.ShouldBind(&subject)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	if subject.Id == 0 { // 新增学科
		err = e.subjectDao.AddSub(&subject)
		if response.CheckError(err, "Add subject error") {
			return response.ResponseOperateFailed()
		}
	}
	//else { // 更新博客
	//	err = e.subjectDao.UpdateSub(&subject)
	//	if response.CheckError(err, "Update subject error") {
	//		return response.ResponseOperateFailed()
	//	}
	//}

	return response.ResponseOperateSuccess()
}

func (e *SubjectController) DeleteSubject(ctx *gin.Context) *response.Response {
	err := e.subjectDao.DeleteSub(ctx.Query("name"))
	if response.CheckError(err, "Delete subject error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseDeleteSuccess()
}

//章节部分的后台管理
type ChapterController struct {
	chapterDao *dao.ChapterDao
}

func NewChapterRouter() *ChapterController {
	return &ChapterController{
		chapterDao: dao.NewChapterDao(),
	}
}

func (e *ChapterController) ChapterList(ctx *gin.Context) *response.Response {
	name := ctx.Query("name")
	//pageNum := utils.QueryInt(ctx, "pageNum")
	//pageSize := utils.QueryInt(ctx, "pageSize")
	chapters, err := e.chapterDao.FindAllChapter(name)
	if response.CheckError(err, "Get Chapter List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(chapters)
}

func (e *ChapterController) AddChapter(ctx *gin.Context) *response.Response {
	var chapter model.ChapterVue
	err := ctx.ShouldBind(&chapter)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.chapterDao.AddChapter(chapter.Subject, chapter.Chapter, chapter.Description)
	if response.CheckError(err, "Add chapter error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}

func (e *ChapterController) DeleteChapter(ctx *gin.Context) *response.Response {
	err := e.chapterDao.DeleteChapter(ctx.Query("subjectName"), ctx.Query("chapterName"))
	if response.CheckError(err, "Delete chapter error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseDeleteSuccess()

}

//题目部分的后台管理
type ExamController struct {
	examDao *dao.ExamDao
}

func NewExamRouter() *ExamController {
	return &ExamController{
		examDao: dao.NewExamDao(),
	}
}

func (e *ExamController) ExamList(ctx *gin.Context) *response.Response {
	a := ctx.Query("typename")
	b := ctx.Query("title")
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")

	exams, err := e.examDao.FindExam(a, b, pageNum, pageSize)
	if response.CheckError(err, "Get Exam List") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(exams)

}

func (e *ExamController) DeleteExam(ctx *gin.Context) *response.Response {
	s := ctx.Query("id")
	id, _ := strconv.Atoi(s)
	err := e.examDao.DeleteExam(id)
	if response.CheckError(err, "Delete exam error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseDeleteSuccess()
}

func (e *ExamController) CreateExam(ctx *gin.Context) *response.Response {
	var exam model.Exam
	err := ctx.ShouldBind(&exam)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.examDao.CreateExam(&exam)
	if response.CheckError(err, "Add Exam error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()

}

func (e *ExamController) UpdateExam(ctx *gin.Context) *response.Response {
	var exam model.Exam
	err := ctx.ShouldBind(&exam)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = e.examDao.UpdateExam(&exam)
	if response.CheckError(err, "update Exam error") {
		return response.ResponseOperateFailed()
	}
	return response.ResponseOperateSuccess()
}
