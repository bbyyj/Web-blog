package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	musics, count, err := m.musicService.GetLimited(pageNum, pageSize)
	if response.CheckError(err, "Get musics error") {
		return response.ResponseQueryFailed()
	}

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

	name := ctx.Query("name")
	println(name)
	if name == "" {
		return response.ResponseDeleteFailed()
	}

	err := m.musicService.DeleteMusic(name)
	if response.CheckError(err, "Delete music error") {
		return response.ResponseDeleteFailed()
	}

	return response.ResponseDeleteSuccess()
}

func (m *MusicController) GetAllMusic(ctx *gin.Context) *response.Response {
	musics, err := m.musicService.GetAll()
	if response.CheckError(err, "Get AllMusic error") {
		return response.ResponseQueryFailed()
	}
	return response.ResponseQuerySuccess(musics)
}
