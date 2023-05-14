package service

import (
	"blog_web/db/dao"
	"blog_web/model"
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

func (m *MusicService) GetLimited(pageNum, pageSize int) ([]model.Music, error) {
	pageStart := (pageNum - 1) * pageSize
	return m.musicDao.FindLimited(pageStart, pageSize)
}

func (m *MusicService) GetCount() (int, error) {
	return m.musicDao.FindTotalCount()
}

func (m *MusicService) AddMusic(music *model.Music) error {
	return m.musicDao.Add(music)
}

func (m *MusicService) DeleteMusic(id int) error {
	return m.musicDao.Delete(id)
}
