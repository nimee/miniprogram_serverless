package dao

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/album"
)

// const tableName = "Counters"
const tableAlbum = "album"

const tablePhoto = "photo"

/*
// ClearCounter 清除Counter
func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Delete(&model.CounterModel{Id: id}).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(tableName).Save(counter).Error
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}

*/

// ClearAlbumer 清除Albumer
func (imp *AlbumerInterfaceImp) ClearAlbumer(theme string) error {
	cli := db.Get()
	if err := cli.Table(tablePhoto).Where("theme = ?", theme).Delete(&album.Photo{}).Error; err != nil {
		fmt.Println(err)
	}
	return cli.Table(tableAlbum).Where("theme = ?", theme).Delete(&album.Album{}).Error //Theme: theme
}

// UpsertAlbumer 更新/写入Albumer
func (imp *AlbumerInterfaceImp) UpsertAlbumer(alblum *album.Album) error {
	cli := db.Get()
	return cli.Table(tableAlbum).Save(alblum).Error
}

// GetAlbumer 查询Albumer
func (imp *AlbumerInterfaceImp) GetAlbumers(groupName string) (*[]album.Album, error) {
	var err error
	var albumersList = new([]album.Album)
	cli := db.Get()
	err = cli.Table(tableAlbum).Where("groupName = ?", groupName).Find(albumersList).Error
	return albumersList, err
}

// ClearPhoto 清除Albumer中的photo
func (imp *AlbumerInterfaceImp) ClearPhoto(photoId string) error {
	cli := db.Get()
	return cli.Table(tablePhoto).Where("photoId = ?", photoId).Delete(&album.Photo{}).Error
}

// UpsertPhoto 更新/写入Albumer中的photo
func (imp *AlbumerInterfaceImp) UpsertPhoto(photo *album.Photo) error {
	cli := db.Get()
	return cli.Table(tablePhoto).Save(photo).Error
}

// GetPhoto 查询Albumer中的photo
func (imp *AlbumerInterfaceImp) GetPhotos(theme string) (*[]album.Photo, error) {
	var err error
	var photoList = new([]album.Photo)
	cli := db.Get()
	err = cli.Table(tablePhoto).Where("theme = ?", theme).Find(photoList).Error
	return photoList, err
}
