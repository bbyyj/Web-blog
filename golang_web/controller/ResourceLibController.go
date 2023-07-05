package controller

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ResourceLibController struct {
	linkService *service.LinkService
}

func NewResourceLibRouter() *ResourceLibController {
	return &ResourceLibController{
		linkService: service.NewLinkService(),
	}
}

func (r *ResourceLibController) LinkList(ctx *gin.Context) *response.Response {
	links, err := r.linkService.GetAllLinks()
	if response.CheckError(err, "Get links error") {
		return response.RQueryFailed()
	}
	categories, err := r.linkService.GetAllCategory()

	if response.CheckError(err, "Get categories error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(links, categories)
}

func (r *ResourceLibController) GetALLResource(ctx *gin.Context) *response.Response {
	//pagestart := utils.StrconvAtoiParm(ctx, "pagestart")
	//pagesize := utils.StrconvAtoiParm(ctx, "pagesize")
	pagestart := utils.DefaultQueryInt(ctx, "pagenum", "1")
	pagesize := utils.DefaultQueryInt(ctx, "pagesize", "10")
	//resource, err := r.linkService.GetAllResource(pagestart, pagesize)
	resource, count, err := r.linkService.GetLimitedResource(pagestart, pagesize)
	fmt.Println(count)
	if response.CheckError(err, "Get AllResource error") {
		return response.RGetResourceFailed()
	}
	return response.RGetResourceSuccess(resource, count)
}

func (r *ResourceLibController) GetALLResourceByCategoryId(ctx *gin.Context) *response.Response {
	//categoryId := utils.StrconvAtoiParm(ctx, "categoryId")
	//pageStart := utils.StrconvAtoiParm(ctx, "pageStart")
	//pageSize := utils.StrconvAtoiParm(ctx, "pageSize")
	categoryId := utils.DefaultQueryInt(ctx, "categoryid", "1")
	pageStart := utils.DefaultQueryInt(ctx, "pagenum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pagesize", "10")
	resource, count, err := r.linkService.GetAllResourceByCategoryId(categoryId, pageStart, pageSize)
	if response.CheckError(err, "Get AllResourceByCategoryId error") {
		return response.RGetResourceFailed()
	}
	return response.RGetResourceSuccess(resource, count)
}

func (r *ResourceLibController) GetAllResourceLikeName(ctx *gin.Context) *response.Response {
	//name := ctx.Param("name")
	name := ctx.Query("name")
	pagenum := utils.QueryInt(ctx, "pagenum")
	pagesize := utils.QueryInt(ctx, "pagesize")
	name = ("%" + name + "%")
	resource, count, err := r.linkService.GetAllResourceLikeName(name, pagenum, pagesize)
	if response.CheckError(err, "Get AllResourceLikeName error") {
		return response.RGetResourceFailed()
	}
	return response.RGetResourceSuccess(resource, count)

}

func (r *ResourceLibController) DownloadReasourceService(ctx *gin.Context) {
	//name := ctx.Param("name")
	//fmt.Println("准备下载，来拿名字")
	name := ctx.Query("name")
	//fmt.Printf("名字拿到咯，是：%s\n来拿地址", name)
	url, err := r.linkService.GetDownloadUrl(name)
	filePath := url.Url
	if response.CheckError(err, "Get DownloadUrl error") {
		ctx.JSON(http.StatusOK, model.ResResult{
			nil, 565, "获取资源失败",
		})
	} else {
		//fmt.Printf("地址拿到咯，是：%s\n准备更新下载次数", filePath)
		err = r.linkService.UpdateResourceDownloadNumByName(name)
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			ctx.JSON(http.StatusNotFound, "File not found")
			return
		}

		if response.CheckError(err, "Get DownloadUrl error") {
			ctx.JSON(http.StatusOK, model.ResResult{
				nil, 565, "获取资源失败",
			})
		}
		//具体下载的实现逻辑
		ctx.Writer.Header().Set("Content-Disposition", "attachment; filename="+name)
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		ctx.File(filePath)
	}

}

func (r *ResourceLibController) GetResourceDetailByName(ctx *gin.Context) *response.Response {
	//name := ctx.Param("name")
	name := ctx.Query("name")
	resource, err := r.linkService.GetAllResourceByName(name)
	if response.CheckError(err, "Get AllResourceLikeName error") {
		return response.RGetResourceFailed()
	}
	return response.RGetResourceSuccess(resource)

}

func (r *ResourceLibController) GetALLCategory(ctx *gin.Context) *response.Response {
	category, err := r.linkService.GetAllCategoryNew()
	if response.CheckError(err, "Get AllResource error") {
		return response.RQueryFailed()
	}
	return response.RQuerySuccess(category)
}

func (l *ResourceLibController) AddReasourceCheckService(ctx *gin.Context) *response.Response {

	var resource model.ResourceManage
	err := ctx.ShouldBind(&resource)
	if response.CheckError(err, "Bind param error") {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
	fmt.Println(resource)
	err = l.linkService.CheckResourceName(resource.ID, resource.Name)
	err = l.linkService.CheckResourceCheckName(resource.ID, resource.Name)
	if response.CheckError(err, "find same name") {
		fmt.Println(err)
		return response.ROperateFailed()
	}
	resource.DownloadNum = 0
	//resource.Url = "D:/Go project/Go_UPandDownload/downloads/" + resource.Name
	resource.FileSize, err = utils.GetFileSize(resource.Url)

	err = l.linkService.AddResourceCheck(&resource)
	if response.CheckError(err, "Add ResourceCheck error") {
		fmt.Println(err)
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}

func (l *ResourceLibController) UploadReasourceCheckService(ctx *gin.Context) {

	file, err := ctx.FormFile("f1")
	if response.CheckError(err, "Get Upload resource error") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	now := time.Now().Unix()
	fp, suf := utils.FileSuffixSplit(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", fp, now, suf) // filename: bg_156435453.jpg

	log.Println(filename)
	dst := fmt.Sprintf(model.FileRoot, filename)
	// 上传文件到指定的目录
	err = ctx.SaveUploadedFile(file, dst)
	if response.CheckError(err, "Upload Resource error") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		"url":     dst,
	})
}
