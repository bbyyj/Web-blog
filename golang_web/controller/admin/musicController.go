package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MusicController struct {
	musicService *service.MusicService
}

func NewMusicRouter() *MusicController {
	return &MusicController{
		musicService: service.NewMusicService(),
	}
}

func (m *MusicController) MusicList(ctx *gin.Context) *response.Response {
	pageNum := utils.DefaultQueryInt(ctx, "pageNum", "1")
	pageSize := utils.DefaultQueryInt(ctx, "pageSize", "10")
	musics, err := m.musicService.GetLimited(pageNum, pageSize)
	if response.CheckError(err, "Get musics error") {
		return response.ResponseQueryFailed()
	}
	count, _ := m.musicService.GetCount()

	return response.ResponseQuerySuccess(musics, count)
}

func (m *MusicController) AddMusic(ctx *gin.Context) *response.Response {

	var music model.Music

	err := ctx.ShouldBind(&music)

	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = m.musicService.AddMusic(&music)

	if response.CheckError(err, "Add music error") {
		return response.ResponseOperateFailed()
	}

	return response.ResponseOperateSuccess()
}

func (m *MusicController) DeleteMusic(ctx *gin.Context) *response.Response {

	idStr := ctx.Query("id")
	id, err0 := strconv.Atoi(idStr)
	if err0 != nil {
		println(err0)
		return response.ResponseDeleteFailed()
	}

	err := m.musicService.DeleteMusic(id)
	if response.CheckError(err, "Delete music error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}
