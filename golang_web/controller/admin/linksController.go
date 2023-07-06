package admin

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
	"time"
)

type LinksController struct {
	linkService *service.LinkService
}

func NewLinksRouter() *LinksController {
	return &LinksController{
		linkService: service.NewLinkService(),
	}
}

func (l *LinksController) LinksList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")

	links, err := l.linkService.GetLimitedLinks(pageNum, pageSize)
	if response.CheckError(err, "Get links error") {
		return response.RQueryFailed()
	}
	categories, err := l.linkService.GetAllCategory()
	if response.CheckError(err, "Get categories error") {
		return response.RQueryFailed()
	}
	count, _ := l.linkService.GetLinkCount()
	return response.RQuerySuccess(links, categories, count)
}

func (l *LinksController) DeleteLink(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteLink(id)
	if response.CheckError(err, "Delete link error") {
		return response.RDeleteFailed()
	}
	return response.RDeleteSuccess()
}

func (l *LinksController) UpdateLink(ctx *gin.Context) *response.Response {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = l.linkService.UpdateLink(&link)
	if response.CheckError(err, "Update link error") {
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}

func (l *LinksController) AddLink(ctx *gin.Context) *response.Response {
	var link model.Link
	err := ctx.ShouldBind(&link)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = l.linkService.AddLink(&link)
	if response.CheckError(err, "Add link error") {
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}

func (l *LinksController) Categories(ctx *gin.Context) *response.Response {
	categories, err := l.linkService.GetAllCategory()
	if response.CheckError(err, "Get categories error") {
		return response.RQueryFailed()
	}
	return response.RQuerySuccess(categories)
}

func (l *LinksController) DeleteCategory(ctx *gin.Context) *response.Response {
	id := utils.QueryInt(ctx, "id")
	err := l.linkService.DeleteCategory(id)
	if response.CheckError(err, "Delete category error") {
		return response.RDeleteFailed()
	}
	return response.RDeleteSuccess()
}

func (l *LinksController) UpdateCategory(ctx *gin.Context) *response.Response {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = l.linkService.UpdateCategory(&category)
	if response.CheckError(err, "Update category error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (l *LinksController) AddCategory(ctx *gin.Context) *response.Response {
	var category model.LinkCategory
	err := ctx.ShouldBind(&category)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	err = l.linkService.AddCategory(&category)
	if response.CheckError(err, "Add category error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (l *LinksController) ResourceList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pagenum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pagesize", "10")
	fmt.Println(pageNum)
	fmt.Println(pageSize)
	links, count, err := l.linkService.GetLimitedResource(pageNum, pageSize)
	if response.CheckError(err, "Get Resource error") {
		return response.RGetResourceFailed()
	}
	categories, err := l.linkService.GetAllCategory()
	return response.RGetResourceSuccess(links, categories, count)
}

func (l *LinksController) ResourceCheckList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pagenum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pagesize", "10")
	fmt.Println(pageNum)
	fmt.Println(pageSize)
	links, count, err := l.linkService.GetLimitedResourceCheck(pageNum, pageSize)
	if response.CheckError(err, "Get Resource error") {
		return response.RGetResourceFailed()
	}
	categories, err := l.linkService.GetAllCategory()
	return response.RGetResourceSuccess(links, categories, count)
}
func (l *LinksController) UpdateResource(ctx *gin.Context) *response.Response {
	var resource model.ResourceManage
	err := ctx.ShouldBind(&resource)
	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	if resource.Url != "" {
		r := l.linkService.GetResourceUrl(resource.ID)
		os.Remove(r.Url)
	}
	if resource.Name != "" {
		err = l.linkService.CheckResourceName(resource.ID, resource.Name)
		if response.CheckError(err, "change name failed") {
			return response.ROperateFailed()
		}
	}
	err = l.linkService.UpdateResource(&resource)
	if response.CheckError(err, "Update link error") {
		//fmt.Println("更新失败")
		return response.ROperateFailed()
	}
	//fmt.Println("更新成功")
	return response.ROperateSuccess()
}

func (l *LinksController) UploadResource(ctx *gin.Context) {
	file, err := ctx.FormFile("f1")
	if response.CheckError(err, "Get Upload resource error") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	now := time.Now().Unix()
	fp, suf := utils.FileSuffixSplit(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", fp, now, suf) // filename: fn_unix.zip
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

func (l *LinksController) AddResource(ctx *gin.Context) *response.Response {

	var resource model.ResourceManage
	err := ctx.ShouldBind(&resource)
	if response.CheckError(err, "Bind param error") {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
	fmt.Println(resource)
	err = l.linkService.CheckResourceName(resource.ID, resource.Name)
	if response.CheckError(err, "find same name") {
		fmt.Println(err)
		return response.ROperateFailed()
	}
	resource.DownloadNum = 0
	resource.FileSize, err = utils.GetFileSize(resource.Url)

	err = l.linkService.AddResource(&resource)
	if response.CheckError(err, "Add Resource error") {
		fmt.Println(err)
		return response.ROperateFailed()
	}
	return response.ROperateSuccess()
}

func (l *LinksController) DeleteResource(ctx *gin.Context) *response.Response {
	Id := utils.QueryInt(ctx, "id")
	id := uint(Id)
	url := l.linkService.GetResourceUrl(id)
	err := os.Remove(url.Url)
	err = l.linkService.DeleteResource(id)
	if response.CheckError(err, "Get Resource error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (l *LinksController) GetResourceLikeName(ctx *gin.Context) *response.Response {
	name := ctx.Query("name")
	pagenum := utils.QueryInt(ctx, "pagenum")
	pagesize := utils.QueryInt(ctx, "pagesize")
	resource, count, err := l.linkService.GetResourceLikeName(name, pagenum, pagesize)
	if response.CheckError(err, "Get Resource error") {
		return response.RGetResourceFailed()
	}

	return response.RGetResourceSuccess(resource, count)
}

func (l *LinksController) ReUploadResource(ctx *gin.Context) {
	name := ctx.Query("name")
	file, err := ctx.FormFile("f1")
	if response.CheckError(err, "Get Upload resource error") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	resource, err := l.linkService.GetAllResourceByName(name)
	if response.CheckError(err, "find resource by name error") {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(name)
	// 上传文件到指定的目录
	err = ctx.SaveUploadedFile(file, resource.Url)

	if response.CheckError(err, "Upload Resource error") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = l.linkService.ReUploadUpdateTime(name)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' ReUploaded!", name),
	})
}

func (l *LinksController) CheckSucceededAddToResource(ctx *gin.Context) *response.Response {
	type RequestBody struct {
		ID int `json:"id"`
	}
	var requestBody RequestBody
	err := ctx.ShouldBind(&requestBody)
	id := uint(requestBody.ID)
	c := l.linkService.GetCheckResourceById(id)
	fmt.Println(c)
	if c.ID == 0 {
		return response.ROperateFailed()
	}
	c.ID = 0
	err = l.linkService.AddResource(c)
	if response.CheckError(err, "add resource error") {
		return response.ROperateFailed()
	}
	err = l.linkService.DeleteResourceCheck(id)
	if response.CheckError(err, "delete resourcecheck error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (l *LinksController) CheckFailedResource(ctx *gin.Context) *response.Response {
	Id := utils.QueryInt(ctx, "id")
	id := uint(Id)
	c := l.linkService.GetCheckResourceById(id)
	if c.ID == 0 {
		return response.ROperateFailed()
	}
	err := os.Remove(c.Url)
	err = l.linkService.DeleteResourceCheck(id)
	if response.CheckError(err, "Delete Check Resource error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}
