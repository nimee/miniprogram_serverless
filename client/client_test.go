package client

import (
	"testing"
	"time"
)

func TestGetAlbums(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get request",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAlbums()
		})
	}
}

func TestUpsertAlbum(t *testing.T) {
	type args struct {
		album *Album
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "upsert album",
			args: args{
				album: &Album{
					Theme:         "王者荣耀",
					Count:         10,
					Photographer:  "Liang",
					Mua:           "Wu",
					T:             time.Now(),
					Location:      "HangZhou",
					GroupName:     "王者荣耀",
					CoverId:       "10",
					StorageType:   "cos",
					StorageParams: []byte{},
					Enable:        true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpsertAlbum(tt.args.album)
		})
	}
}

func TestDeleteAlbum(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete album",
			args: args{
				id: 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteAlbum(tt.args.id)
		})
	}
}

func TestGetPhoto(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "query photos on a themeid",
			args: args{
				id: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPhoto(tt.args.id)
		})
	}
}

func TestUpsertPhoto(t *testing.T) {
	type args struct {
		photo *Photo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert photo",
			args: args{
				photo: &Photo{
					AlbumId: 12,
					PhotoId: "ccc",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpsertPhoto(tt.args.photo)
		})
	}
}

func TestDeletePhoto(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete photo id",
			args: args{
				id: "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeletePhoto(tt.args.id)
		})
	}
}
