package dao

import (
	"wxcloudrun-golang/db/album"
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounter(id int32) (*model.CounterModel, error)
	UpsertCounter(counter *model.CounterModel) error
	ClearCounter(id int32) error
}

// CounterInterfaceImp 计数器数据模型实现
//type CounterInterfaceImp struct{}

// Imp 实现实例
//var Imp CounterInterface = &CounterInterfaceImp{}

type AlbumerInterface interface {
	ClearAlbumer(id uint) error
	UpsertAlbumer(alblum *album.Album) error
	GetAlbumers() (*[]album.Album, error)

	ClearPhoto(photoId string) error
	UpsertPhoto(photo *album.Photo) error
	GetPhotos(id uint) (*[]album.Photo, error)
}

type AlbumerInterfaceImp struct{}

var Imp AlbumerInterface = &AlbumerInterfaceImp{}
