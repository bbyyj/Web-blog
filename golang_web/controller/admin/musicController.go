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

	messages, count, err := m.musicService.MusicList(pageNum, pageSize)

	if response.CheckError(err, "Get MusicList error") {
		return response.RQueryFailed()
	}

	return response.RQuerySuccess(messages, count)
}

func (m *MusicController) AddMusic(ctx *gin.Context) *response.Response {

	var music model.Music

	err := ctx.ShouldBind(&music)
	println(music.Url)
	println(music.Name)

	if response.CheckError(err, "Bind param error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}

	err = m.musicService.AddMusic(&music)

	if response.CheckError(err, "Add music error") {
		return response.ROperateFailed()
	}

	return response.ROperateSuccess()
}

func (m *MusicController) DeleteMusic(ctx *gin.Context) *response.Response {

	id := utils.QueryInt(ctx, "id")
	println(id)

	err := m.musicService.DeleteMusic(id)
	if response.CheckError(err, "Delete music error") {
		return response.RDeleteFailed()
	}

	return response.RDeleteSuccess()
}

func (m *MusicController) GetAllMusic(ctx *gin.Context) *response.Response {
	musics, err := m.musicService.GetAll()
	if response.CheckError(err, "Get AllMusic error") {
		return response.RQueryFailed()
	}
	return response.RQuerySuccess(musics)
}
