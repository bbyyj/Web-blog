package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	"testing"
)

func TestFindAll(t *testing.T) {
	music := dao.NewMusicDao()
	musics, err := music.FindAll()
	if err != nil {
		t.Fatal("Find All Error")
	}
	t.Logf("Find All Success, musics = %v", musics)
}

func TestMusicList(t *testing.T) {
	music := dao.NewMusicDao()
	musics, err := music.MusicList(1, 10)
	if err != nil {
		t.Fatal("Music List Error")
	}
	t.Logf("Music List Success, musics = %v", musics)
}

func TestFindTotalCount(t *testing.T) {
	music := dao.NewMusicDao()
	count, err := music.FindTotalCount()
	if err != nil {
		t.Fatal("Find Total Count Error")
	}
	t.Logf("Find Total Count Success, count = %v", count)
}

//	func (m *MusicDao) Add(music *model.Music) error {
//		_, err := sqldb.Exec(m.sql[3], music.Name, music.Artist, music.Url, music.Cover)
//		return err
//	}
//
//	func (m *MusicDao) Delete(id int) error {
//		_, err := sqldb.Exec(m.sql[4], id)
//		return err
//	}
func TestAdd(t *testing.T) {
	music := dao.NewMusicDao()
	var newMusic model.Music
	newMusic.Name = "Everything's Alright"
	newMusic.Artist = "塞壬唱片-MSR,DJ Okawari,澤田かおり"
	newMusic.Url = "http://music.163.com/song/media/outer/url?id=1460626792.mp3"
	newMusic.Cover = "http://p1.music.126.net/C_dc-WuRU4YKJREo4A2Wbw==/109951165115215198.jpg?param=130y130"
	err := music.Add(&newMusic)
	if err != nil {
		t.Fatal("Add Error")
	}
	t.Log("Add Success")
}

func TestDelete(t *testing.T) {
	music := dao.NewMusicDao()
	maxId, err := music.GetMaxMusicId()
	err = music.Delete(maxId)
	if err != nil {
		t.Fatal("DeleteError")
	}
	t.Log("Delete Success")
}
