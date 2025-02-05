package dao

import (
	"os"
	"reflect"
	"testing"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/album"
)

func init() {
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "root")
	os.Setenv("MYSQL_ADDRESS", "192.168.129.128:3306")
	db.Init()
}

func TestAlbumerInterfaceImp_UpsertAlbumer(t *testing.T) {

	type args struct {
		alblum *album.Album
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		wantErr bool
	}{
		{
			name: "upsert test",
			imp:  Imp,
			args: args{
				alblum: &album.Album{
					ID:            1,
					Theme:         "云栖竹径",
					Count:         10,
					Photographer:  "Liang",
					Mua:           "Wu",
					T:             time.Now(),
					Location:      "HangZhou",
					GroupName:     "云栖竹径",
					CoverId:       "0",
					StorageType:   "cos",
					StorageParams: []byte{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imp.UpsertAlbumer(tt.args.alblum); (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.UpsertAlbumer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlbumerInterfaceImp_ClearAlbumer(t *testing.T) {

	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		wantErr bool
	}{
		{
			name: "delete test",
			imp:  Imp,
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imp.ClearAlbumer(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.ClearAlbumer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlbumerInterfaceImp_GetAlbumers(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		want    *album.Album
		wantErr bool
	}{
		{
			name:    "query test",
			imp:     Imp,
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imp.GetAlbumers()
			if (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.GetAlbumers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AlbumerInterfaceImp.GetAlbumers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbumerInterfaceImp_UpsertPhoto(t *testing.T) {
	type args struct {
		photo *album.Photo
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		wantErr bool
	}{
		{
			name: "insert photo",
			imp:  Imp,
			args: args{
				photo: &album.Photo{
					AlbumId: 1,
					PhotoId: "b",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imp.UpsertPhoto(tt.args.photo); (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.UpsertPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlbumerInterfaceImp_ClearPhoto(t *testing.T) {
	type args struct {
		photoId string
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		wantErr bool
	}{
		{
			name: "delete photo",
			imp:  Imp,
			args: args{
				photoId: "c",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.imp.ClearPhoto(tt.args.photoId); (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.ClearPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlbumerInterfaceImp_GetPhotos(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		imp     AlbumerInterface
		args    args
		want    *[]album.Photo
		wantErr bool
	}{
		{
			name: "query photos",
			imp:  Imp,
			args: args{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imp.GetPhotos(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlbumerInterfaceImp.GetPhotos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AlbumerInterfaceImp.GetPhotos() = %v, want %v", got, tt.want)
			}
		})
	}
}
