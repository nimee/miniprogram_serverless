package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Body struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type Album struct {
	ID            uint      `json:"id"`
	Theme         string    `gorm:"index;primaryKey;column:theme;type:varchar(255)" json:"theme"`
	Count         int16     `gorm:"column:count;type:int(11)" json:"count"`
	Photographer  string    `gorm:"column:photographer;type:varchar(255)" json:"photographer"`
	Mua           string    `gorm:"column:mua;type:varchar(255)" json:"mua"`
	T             time.Time `gorm:"type:timestamp;column:t" json:"t"`
	Location      string    `gorm:"column:location;type:varchar(255)" json:"location"`
	GroupName     string    `gorm:"column:groupName;type:varchar(255)" json:"groupName"`
	CoverId       string    `gorm:"column:coverId;type:varchar(255)" json:"coverId"`
	StorageType   string    `gorm:"column:storageType;type:varchar(255)" json:"storageType"`
	StorageParams []byte    `gorm:"column:storageParams;type:varchar(255)" json:"storageParams"`
	Enable        bool      `gorm:"column:enable;type:TINYINT(1);default:1" json:"enable"`
	Note          string    `gorm:"column:note;type:varchar(1024)" json:"note"`
}

type Photo struct {
	AlbumId uint   `json:"albumId"`
	PhotoId string `json:"photoId"`
	Enable  bool   `json:"enable"`
}

func GetAlbums() {
	url := "http://127.0.0.1:80/api/album?action=query"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//fmt.Println(resp)

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	defer resp.Body.Close()

	// print response body
	//fmt.Println(string(body))
	var b Body
	err = json.Unmarshal(body, &b)
	if err != nil {
		fmt.Println("parse data err")
	}
	var albums []Album
	json.Unmarshal(b.Data, &albums)
	fmt.Println(albums)
}

func GetPhoto(id uint) {
	url := fmt.Sprintf("http://127.0.0.1:80/api/photo?action=query&albumId=%d", id)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//fmt.Println(resp)

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	resp.Body.Close()

	// print response body
	fmt.Println(string(body))
	var b Body
	err = json.Unmarshal(body, &b)
	if err != nil {
		fmt.Println("parse data err")
	}
	var photos []Photo
	json.Unmarshal(b.Data, &photos)
	fmt.Println(photos)
}

func UpsertAlbum(album *Album) {
	url := "http://127.0.0.1:80/api/album?action=update"
	a, _ := json.Marshal(album)

	requst, err := http.NewRequest("POST", url, bytes.NewBuffer(a))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(requst)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res)
}

func UpsertPhoto(photo *Photo) {
	url := "http://127.0.0.1:80/api/photo?action=update"
	p, _ := json.Marshal(photo)

	requst, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(requst)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println(res)
}

func DeleteAlbum(id uint) {
	url := fmt.Sprintf("http://127.0.0.1:80/api/album?action=delete&albumId=%d", id)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	resp.Body.Close()

	fmt.Println(string(body))

}

func DeletePhoto(id string) {
	url := fmt.Sprintf("http://127.0.0.1:80/api/photo?action=delete&photoId=%s", id)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	resp.Body.Close()

	fmt.Println(string(body))

}
