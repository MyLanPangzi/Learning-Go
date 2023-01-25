package impl

import (
	"database/sql"
	"web-service-gin/entity"
)

type AlbumServiceImpl struct {
	db *sql.DB
}

func NewAlbumService(db *sql.DB) *AlbumServiceImpl {
	return &AlbumServiceImpl{db: db}
}

func (a *AlbumServiceImpl) Add(album entity.Album) (entity.Album, error) {
	result, err := a.db.Exec("insert into album (title, artist, price) values (?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return entity.Album{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entity.Album{}, err
	}
	album.ID = id
	return album, nil
}

func (a *AlbumServiceImpl) Get(ID int64) (entity.Album, error) {
	row := a.db.QueryRow("select id, title, artist, price from album where id = ? ", ID)
	r := entity.Album{}
	err := row.Scan(&r.ID, &r.Title, &r.Artist, &r.Price)
	if err != nil {
		return entity.Album{}, err
	}
	return r, nil
}

func (a *AlbumServiceImpl) GetAll() []entity.Album {
	rows, err := a.db.Query("select id, title, artist, price from album")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var r []entity.Album
	for rows.Next() {
		var album entity.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		r = append(r, album)
		if err != nil {
			return nil
		}
	}
	return r
}
