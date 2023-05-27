package service

import (
	"blog_web/db/dao"
	"blog_web/model"
	"blog_web/utils"
)

type MusicService struct {
	musicDao *dao.MusicDao
}

func NewMusicService() *MusicService {
	return &MusicService{
		musicDao: dao.NewMusicDao(),
	}
}

func (m *MusicService) GetAll() ([]model.Music, error) {
	return m.musicDao.FindAll()
}

func (m *MusicService) GetLimited(pageNum, pageSize int) ([]model.Music, int, error) {
	pageStart := (pageNum - 1) * pageSize
	musics, err := m.musicDao.FindLimited(pageStart, pageSize)
	if err != nil {
		utils.Logger().Warning("Get musics error:%v", err)
		return nil, 0, err
	}
	count, _ := m.musicDao.FindTotalCount()
	return musics, count, nil
}

func (m *MusicService) GetCount() (int, error) {
	return m.musicDao.FindTotalCount()
}

func (m *MusicService) AddMusic(music *model.Music) error {
	return m.musicDao.Add(music)
}

func (m *MusicService) DeleteMusic(name string) error {
	return m.musicDao.Delete(name)
}
