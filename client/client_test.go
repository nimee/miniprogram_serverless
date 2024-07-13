package client

import (
	"testing"
	"time"
)

func TestGetAlbum(t *testing.T) {
	type args struct {
		group string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get request",
			args: args{group: "nature"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAlbum(tt.args.group)
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
					Theme:         "lotus",
					Count:         10,
					Photographer:  "Liang",
					Mua:           "Wu",
					T:             time.Now(),
					Location:      "HangZhou",
					GroupName:     "nature",
					CoverId:       "0",
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
		theme string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete album",
			args: args{
				theme: "two",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteAlbum(tt.args.theme)
		})
	}
}

func TestGetPhoto(t *testing.T) {
	type args struct {
		theme string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "query photos on a theme",
			args: args{
				theme: "lotus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPhoto(tt.args.theme)
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
					Theme:   "bamboo",
					PhotoId: "1",
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
