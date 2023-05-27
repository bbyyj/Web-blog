package dao

import (
	"blog_web/model"
)

type MusicDao struct {
	sql []string
}

func NewMusicDao() *MusicDao {
	return &MusicDao{
		sql: []string{
			`SELECT id, name, artist, url, cover FROM t_music;`,
			`SELECT id, name, artist, url, cover FROM t_music ORDER BY artist LIMIT ?, ? ;`,
			`SELECT COUNT(*) FROM t_music`,
			`INSERT INTO t_music (name, artist, url, cover) VALUES (?, ?, ?, ?);`,
			`DELETE FROM t_music WHERE name = ?;`,
		},
	}
}

func (m *MusicDao) FindAll() (musics []model.Music, err error) {
	err = sqldb.Select(&musics, m.sql[0])
	return
}

func (m *MusicDao) MusicList(pageStart, PageSize int) (msg []model.Music, err error) {
	err = sqldb.Select(&msg, m.sql[1], pageStart, PageSize)
	return
}

func (m *MusicDao) FindTotalCount() (count int, err error) {
	err = sqldb.Get(&count, m.sql[2])
	return
}

func (m *MusicDao) Add(music *model.Music) error {
	_, err := sqldb.Exec(m.sql[3], music.Name, music.Artist, music.Url, music.Cover)
	return err
}

func (m *MusicDao) Delete(name string) error {
	_, err := sqldb.Exec(m.sql[4], name)
	println(err.Error())
	return err
}
