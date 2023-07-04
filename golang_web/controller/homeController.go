package controller

import (
	"blog_web/db/service"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	blogService *service.BlogService
	userService *service.UserService
}

func NewHomeRouter() *HomeController {
	return &HomeController{
		blogService: service.NewBlogService(),
		userService: service.NewUserService(),
	}
}

// 主页博客展示
func (h *HomeController) HomeListBlogs(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	utils.Logger().Debug("pageNum:%v, pageSize:%v", pageNum, pageSize)

	if pageNum <= 0 || pageSize <= 0 {
		return response.ResponseQueryFailed()
	}
	blogs, err := h.blogService.GetHomePageBlogs(pageNum, pageSize)
	if response.CheckError(err, "Get Blogs error") {
		return response.ResponseQueryFailed()
	}
	count, err := h.blogService.GetBolgCount()
	if response.CheckError(err, "Get Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs, count)
}

func (h *HomeController) GetHomePageUInfo(ctx *gin.Context) *response.Response {
	user, tagn, err := h.userService.GetInfo()
	if response.CheckError(err, "Get Userinfo error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(user, tagn)
}

// 浏览博客详情
func (h *HomeController) GetDetailedBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	blog, tags, err := h.blogService.GetDetailedBlog(id)
	if response.CheckError(err, "Get Detailed blog error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blog, tags)
}

// 搜索博客
func (h *HomeController) SearchBlog(ctx *gin.Context) *response.Response {
	keyWord := ctx.Query("keyWord")
	blogs, err := h.blogService.GetBlogsByKeyWord(keyWord)
	if response.CheckError(err, "Search Blogs error") {
		return response.ResponseQueryFailed()
	}

	return response.ResponseQuerySuccess(blogs)
}
