package router

import (
	"blog_web/controller"
	"blog_web/controller/admin"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	registerBlogRouters(engine)
	registerBlogManageRouter(engine)
}

func registerBlogRouters(engine *gin.Engine) {
	homeRouter := controller.NewHomeRouter()
	blogGroup := engine.Group("/api/myblog")
	{
		blogGroup.GET("/blogLists", Decorate(homeRouter.HomeListBlogs))
		blogGroup.GET("/userInfo", Decorate(homeRouter.GetHomePageUInfo))
		blogGroup.GET("/detailedBlog", Decorate(homeRouter.GetDetailedBlog))
		blogGroup.GET("/search", Decorate(homeRouter.SearchBlog))
		blogGroup.GET("/blogsByTag", Decorate(homeRouter.GetAllBlogByTags))
	}

	typeListRouter := controller.NewTypeListRouter()
	{
		blogGroup.GET("/typeList", Decorate(typeListRouter.GetTypeList))
		blogGroup.GET("/typeBlogList", Decorate(typeListRouter.GetBlogListByTypeid))
	}

	tagListRouter := controller.NewTagListRouter()
	{
		blogGroup.GET("/tagList", Decorate(tagListRouter.GetTagList))
		blogGroup.GET("/tagBlogList", Decorate(tagListRouter.GetBlogListByTagId))
	}

	resourceLibRouter := controller.NewResourceLibRouter()
	{
		blogGroup.GET("/links", Decorate(resourceLibRouter.LinkList))
		blogGroup.GET("/t/category", Decorate(resourceLibRouter.GetALLCategory))
		blogGroup.GET("/t/pageresource", Decorate(resourceLibRouter.GetALLResource))
		blogGroup.GET("/t/pageresourcebycategoryid", Decorate(resourceLibRouter.GetALLResourceByCategoryId))
		blogGroup.GET("/t/queryresource", Decorate(resourceLibRouter.GetAllResourceLikeName))
		blogGroup.GET("/t/downloadresource", resourceLibRouter.DownloadReasourceService)
		blogGroup.POST("/t/uploadresourcecheck", resourceLibRouter.UploadReasourceCheckService)
		blogGroup.POST("/t/addresourcecheck", Decorate(resourceLibRouter.AddReasourceCheckService))
	}

	musicFrontRouter := admin.NewMusicRouter()
	{
		blogGroup.GET("/getAllMusic", Decorate(musicFrontRouter.GetAllMusic))
	}
	subjectRouter := controller.NewSubjectRouter()
	{
		blogGroup.GET("/subjectList", Decorate(subjectRouter.SubjectList))
	}

	chapterRouter := controller.NewChapterRouter()
	{
		blogGroup.GET("/chapterList", Decorate(chapterRouter.ChapterList))
	}

	examRouter := controller.NewExamRouter()
	{
		blogGroup.GET("/examList", Decorate(examRouter.ExamList))
	}

	electionRouter := controller.NewElectionRouter()
	{
		blogGroup.GET("/electionList", Decorate(electionRouter.ElectionList))
		blogGroup.GET("/electionDetailedList", Decorate(electionRouter.ElectionDetailedList))
		blogGroup.GET("/electionListNoClass", Decorate(electionRouter.ElectionListNoClass))
	}

	electionCommentRouter := controller.NewElectionCommentRouter()
	{
		blogGroup.GET("/electionCommentList", Decorate(electionCommentRouter.ElectionCommentList))
		blogGroup.POST("/addElectionComment", Decorate(electionCommentRouter.AddElectionComment))
	}

	askboxFrontRouter := controller.NewAskboxFrontRouter()
	{
		blogGroup.GET("/getAnsweredQA", Decorate(askboxFrontRouter.GetAnsweredQA))
		blogGroup.POST("/addNewQuestion", Decorate(askboxFrontRouter.AddNewQuestion))
		blogGroup.POST("/appendOldQuestion", Decorate(askboxFrontRouter.AppendOldQuestion))
		blogGroup.PUT("/clickLikes", Decorate(askboxFrontRouter.ClickLikes))
	}
}

func registerBlogManageRouter(engine *gin.Engine) {
	loginRouter := admin.NewLoginRouter()
	engine.POST("/api/admin/login", Decorate(loginRouter.Login))

	adminGroup := engine.Group("/api/admin")
	adminGroup.Use(admin.LoginAuthenticationMiddleware())
	blogRouter := admin.NewBlogRouter()
	{
		adminGroup.GET("/searchBlogs", Decorate(blogRouter.SearchBlogs))
		adminGroup.DELETE("/deleteBlog", Decorate(blogRouter.DeleteBlog))
		adminGroup.GET("/getFullBlog", Decorate(blogRouter.GetFullBlog))
		adminGroup.PUT("/updateBlog", Decorate(blogRouter.AddBlog))
	}

	typeRouter := admin.NewTypeRouter()
	{
		adminGroup.GET("/getAllTypes", Decorate(typeRouter.GetAllTypes))
		adminGroup.GET("/getPageTypes", Decorate(typeRouter.GetOnePageTypes))
		adminGroup.GET("/checkTypeExist", Decorate(typeRouter.CheckTypeExist))
		adminGroup.DELETE("/deleteType", Decorate(typeRouter.DeleteType))
		adminGroup.PUT("/updateType", Decorate(typeRouter.UpdateType))
		adminGroup.POST("/addType", Decorate(typeRouter.AddType))
	}

	tagRouter := admin.NewTagRouter()
	{
		adminGroup.GET("/getAllTags", Decorate(tagRouter.GetAllTags))
		adminGroup.GET("/getPageTags", Decorate(tagRouter.GetOnePageTags))
		adminGroup.GET("/checkTagExist", Decorate(tagRouter.CheckTagExist))
		adminGroup.DELETE("/deleteTag", Decorate(tagRouter.DeleteTag))
		adminGroup.PUT("/updateTag", Decorate(tagRouter.UpdateTag))
		adminGroup.POST("/addTag", Decorate(tagRouter.AddTag))
	}

	linksRouter := admin.NewLinksRouter()
	{
		adminGroup.GET("/pageLinks", Decorate(linksRouter.LinksList))
		adminGroup.DELETE("/deleteLink", Decorate(linksRouter.DeleteLink))
		adminGroup.POST("/addLink", Decorate(linksRouter.AddLink))
		adminGroup.PUT("/updateLink", Decorate(linksRouter.UpdateLink))
		adminGroup.GET("/categories", Decorate(linksRouter.Categories))
		adminGroup.DELETE("/deleteCategory", Decorate(linksRouter.DeleteCategory))
		adminGroup.POST("/addCategory", Decorate(linksRouter.AddCategory))
		adminGroup.PUT("/updateCategory", Decorate(linksRouter.UpdateCategory))
		adminGroup.GET("/t/pageresource", Decorate(linksRouter.ResourceList))
		adminGroup.PUT("/t/updateresource", Decorate(linksRouter.UpdateResource))
		adminGroup.POST("/t/uploadresource", linksRouter.UploadResource)
		adminGroup.POST("/t/addresource", Decorate(linksRouter.AddResource))
		adminGroup.DELETE("/t/deleteresource", Decorate(linksRouter.DeleteResource))
		adminGroup.GET("/t/queryresource/", Decorate(linksRouter.GetResourceLikeName))
		adminGroup.POST("/t/reupload", linksRouter.ReUploadResource)
		adminGroup.PUT("/t/checksucceeded", Decorate(linksRouter.CheckSucceededAddToResource))
		adminGroup.DELETE("/t/checkfailed", Decorate(linksRouter.CheckFailedResource))
		adminGroup.GET("/t/pageresourcecheck", Decorate(linksRouter.ResourceCheckList))
	}

	musicRouter := admin.NewMusicRouter()
	{
		adminGroup.GET("/musicList", Decorate(musicRouter.MusicList))
		adminGroup.POST("/addMusic", Decorate(musicRouter.AddMusic))
		adminGroup.DELETE("/deleteMusic", Decorate(musicRouter.DeleteMusic))
	}

	askboxBackRouter := admin.NewAskBoxBackRouter()
	{
		adminGroup.GET("/getAllQA", Decorate(askboxBackRouter.GetAllQA))
		adminGroup.GET("/getUnansweredQA", Decorate(askboxBackRouter.GetUnansweredQA))
		adminGroup.GET("/getAnsweredQAPage", Decorate(askboxBackRouter.GetAnsweredQAPage))
		adminGroup.PUT("/addAnswer", Decorate(askboxBackRouter.AddAnswer))
		adminGroup.PUT("/modifyAnswer", Decorate(askboxBackRouter.ModifyAnswer))
		adminGroup.DELETE("/deleteQuestion", Decorate(askboxBackRouter.DeleteQuestion))
	}
	subjectRouter := admin.NewSubjectRouter()
	{
		adminGroup.POST("/addSubject", Decorate(subjectRouter.AddSubject))
		adminGroup.DELETE("/deleteSubject", Decorate(subjectRouter.DeleteSubject))
	}

	chapterRouter := admin.NewChapterRouter()
	{
		adminGroup.GET("/chapterList", Decorate(chapterRouter.ChapterList))
		adminGroup.POST("/addChapter", Decorate(chapterRouter.AddChapter))
		adminGroup.DELETE("/deleteChapter", Decorate(chapterRouter.DeleteChapter))
	}
	examRouter := admin.NewExamRouter()
	{
		adminGroup.GET("/examList", Decorate(examRouter.ExamList))
		adminGroup.DELETE("/deleteExam", Decorate(examRouter.DeleteExam))
		adminGroup.POST("/createExam", Decorate(examRouter.CreateExam))
		adminGroup.PUT("/updateExam", Decorate(examRouter.UpdateExam))
	}

	electionRouter := admin.NewElectionRouter()
	{
		adminGroup.GET("/electionByClass", Decorate(electionRouter.ElectionByClass))
		adminGroup.GET("/electionList", Decorate(electionRouter.ElectionAllList))
		adminGroup.DELETE("/deleteElection", Decorate(electionRouter.DeleteElection))
		adminGroup.POST("/addElection", Decorate(electionRouter.AddElection))
		adminGroup.PUT("/updateElection", Decorate(electionRouter.UpdateElection))
		adminGroup.GET("/electionNoPage", Decorate(electionRouter.ElectionListNoPage))
	}

	electionCommentRouter := admin.NewElectionCommentRouter()
	{
		adminGroup.GET("/electionCommentList", Decorate(electionCommentRouter.ElectionCommentAllList))
		adminGroup.GET("/electionByClassification", Decorate(electionCommentRouter.ElectionCommentListByClassification))
		adminGroup.DELETE("/deleteElectionComment", Decorate(electionCommentRouter.DeleteElectionComment))
	}
}
