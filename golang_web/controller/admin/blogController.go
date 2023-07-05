package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BlogController struct {
	blogService *service.BlogService
}

func NewBlogRouter() *BlogController {
	return &BlogController{
		blogService: service.NewBlogService(),
	}
}

// 提供三种交叉分类的方法来后台管理文章
func (b *BlogController) SearchBlogs(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "5")
	blogTitle := ctx.Query("blogTitle")
	typeId := utils.QueryInt(ctx, "typeId")
	recommended := ctx.Query("recommended")

	blogs, count, err := b.blogService.GetBlogsByTitleOrTypeOrRecommend(pageNum, pageSize, blogTitle, typeId, recommended)
	if response.CheckError(err, "Search blogs error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(blogs, count)
}

func (b *BlogController) DeleteBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := b.blogService.DeleteBlog(id)
	if response.CheckError(err, "Delete blog error") {
		return response.RDeleteFailed()
	}

	return response.RDeleteSuccess()
}

// 包括文章正文内容的 并使用tag服务里的操作返回一串tag
func (b *BlogController) GetFullBlog(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	blogs, tags, err := b.blogService.GetFullBlog(id)
	if response.CheckError(err, "Get Full blog error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(blogs, tags)
}

func (b *BlogController) AddBlog(ctx *gin.Context) *response.Response {
	var blog model.FullBlog
	err := ctx.ShouldBind(&blog)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	tm := time.Now()
	if blog.Id == 0 {
		blog.Views = 0
		blog.CreateTime = tm
	}
	blog.UpdateTime = tm

	// 新的记录插入前id为0，插入到数据库时自动更新
	if blog.Id == 0 { // 新记录插入
		err = b.blogService.AddBlog(&blog)
		if response.CheckError(err, "Add blog error") {
			return response.ROperateFailed()
		}
	} else { // 旧记录更新
		err = b.blogService.UpdateBlog(&blog)
		if response.CheckError(err, "Update blog error") {
			return response.ROperateFailed()
		}
	}

	return response.ROperateSuccess()
}
