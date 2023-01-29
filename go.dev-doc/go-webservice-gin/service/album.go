package service

import "web-service-gin/entity"

type AlbumService interface {
	Add(album entity.Album) (entity.Album, error)
	Get(ID int64) (entity.Album, error)
	GetAll() []entity.Album
}
