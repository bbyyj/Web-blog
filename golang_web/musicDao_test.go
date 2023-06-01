package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	"testing"
)

func TestFindAll(t *testing.T) {
	music := dao.NewMusicDao()
	_, err := music.FindAll()
	if err != nil {
		println(err.Error())
		t.Fatal("Find All Error")
	}
	t.Log("Find All Success")
}

func TestMusicList(t *testing.T) {
	music := dao.NewMusicDao()
	_, err := music.MusicList(1, 10)
	if err != nil {
		println(err.Error())
		t.Fatal("Music List Error")
	}
	t.Log("Music List Success")
}

func TestFindTotalCount(t *testing.T) {
	music := dao.NewMusicDao()
	_, err := music.FindTotalCount()
	if err != nil {
		println(err.Error())
		t.Fatal("Find Total Count Error")
	}
	t.Log("Find Total Count Success")
}

// 添加并删除同一音乐
func TestAdd(t *testing.T) {
	music := dao.NewMusicDao()
	var newMusic model.Music
	newMusic.Name = "Everything's Alright"
	newMusic.Artist = "塞壬唱片-MSR,DJ Okawari,澤田かおり"
	newMusic.Url = "http://music.163.com/song/media/outer/url?id=1460626792.mp3"
	newMusic.Cover = "http://p1.music.126.net/C_dc-WuRU4YKJREo4A2Wbw==/109951165115215198.jpg?param=130y130"
	err := music.Add(&newMusic)
	if err != nil {
		println(err.Error())
		t.Fatal("Add Error")
	}
	t.Log("Add Success")
}

func TestDelete(t *testing.T) {
	music := dao.NewMusicDao()
	maxId, err := music.GetMaxMusicId()
	err = music.Delete(maxId)
	if err != nil {
		println(err.Error())
		t.Fatal("Delete Error")
	}
	t.Log("Delete Success")
}
