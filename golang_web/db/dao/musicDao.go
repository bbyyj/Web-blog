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
			`SELECT id, name, artist, url, cover FROM music;`,
			`SELECT id, name, artist, url, cover FROM music ORDER BY artist LIMIT ?, ? ;`,
			`SELECT COUNT(*) FROM music`,
			`INSERT INTO music (name, artist, url, cover) VALUES (?, ?, ?, ?);`,
			`DELETE FROM music WHERE id = ?;`,
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

func (m *MusicDao) Delete(id int) error {
	_, err := sqldb.Exec(m.sql[4], id)
	return err
}

func (m *MusicDao) GetMaxMusicId() (maxID int, err error) {
	err = sqldb.Get(&maxID, "SELECT Max(id) FROM music")
	//println(maxID)
	return
}
