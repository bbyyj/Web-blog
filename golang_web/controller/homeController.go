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

// 显示主页的文章 不包含文章内容
func (h *HomeController) HomeListBlogs(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	utils.Logger().Debug("pageNum:%v, pageSize:%v", pageNum, pageSize)

	if pageNum <= 0 || pageSize <= 0 {
		return response.RQueryFailed()
	}
	blogs, err := h.blogService.GetHomePageBlogs(pageNum, pageSize)
	if response.CheckError(err, "Get Blogs error") {
		return response.RQueryFailed()
	}
	count, err := h.blogService.GetBlogCount()
	if response.CheckError(err, "Get Blogs error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(blogs, count)
}

// 使用user服务 显示用户名和头像
func (h *HomeController) GetHomePageUInfo(ctx *gin.Context) *response.Response {
	user, tagn, err := h.userService.GetInfo()
	if response.CheckError(err, "Get Userinfo error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(user, tagn)
}

// 点进具体的一篇文章 包括文章内容 并刷新浏览次数
func (h *HomeController) GetDetailedBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	blog, tags, err := h.blogService.GetDetailedBlog(id)
	if response.CheckError(err, "Get Detailed blog error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(blog, tags)
}

// 搜索标题或正文里含有关键字的文章（模糊匹配）
func (h *HomeController) SearchBlog(ctx *gin.Context) *response.Response {
	keyWord := ctx.Query("keyWord")
	blogs, err := h.blogService.GetBlogsByKeyWord(keyWord)
	if response.CheckError(err, "Search Blogs error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(blogs)
}

// 根据TagId获取分页博客列表
func (h *HomeController) GetAllBlogByTags(ctx *gin.Context) *response.Response {
	pageNum := utils.QueryInt(ctx, "pageNum")
	pageSize := utils.QueryInt(ctx, "pageSize")
	tagId := utils.QueryInt(ctx, "tagId")
	utils.Logger().Debug("pageNum:%v, pageSize:%v", pageNum, pageSize)

	if pageNum <= 0 || pageSize <= 0 {
		return response.RQueryFailed()
	}
	blogs, count, err := h.blogService.GetBlogsByTagId(tagId, pageNum, pageSize)
	if response.CheckError(err, "Get Blogs error") {
		return response.RQueryFailed()
	}
	return response.RQuerySuccess(blogs, count)
}
