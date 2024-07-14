package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"wxcloudrun-golang/db/album"
	"wxcloudrun-golang/db/dao"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getIndex()
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Fprint(w, data)
}

// AlbumHandler album接口
func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {
		action := getAction(r)
		if action == "query" {
			albums, err := getAlbums(r)
			if err != nil {
				res.Code = 404
				res.ErrorMsg = err.Error()
			} else {
				res.Data = albums
			}
		} else if action == "delete" {
			code, err := deleteAlbum(r)
			if err != nil {
				res.Code = code
				res.ErrorMsg = err.Error()
			} else {
				res.Data = code
			}
		}
	} else if r.Method == http.MethodPost {
		code, err := modifyAlbum(r)
		if err != nil {
			res.Code = code
			res.ErrorMsg = err.Error()
		} else {
			res.Data = code
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// modifyAlbum 更新album
func modifyAlbum(r *http.Request) (int, error) {
	action := getAction(r)

	code := 200
	var err error
	if action == "update" {
		code, err = upsertAlbum(r)
		// if err != nil {
		// 	return code, err
		// }
	} else if action == "delete" {
		//code, err = deleteAlbum(r)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
		code = 500
	}

	return code, err
}

// modifyPhoto 更新photo表中的一条记录
func modifyPhoto(r *http.Request) (int, error) {
	action := getAction(r)

	code := 200
	var err error
	if action == "update" {
		code, err = upsertPhoto(r)
		// if err != nil {
		// 	return code, err
		// }
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
		code = 500
	}

	return code, err
}

// upsertAlbum 更新或修改album
func upsertAlbum(r *http.Request) (int, error) {
	album, err := parseAlbum(r)
	if err != nil {
		return 500, err
	}
	err = dao.Imp.UpsertAlbumer(album)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

// upsertPhoto 更新或修改photo
func upsertPhoto(r *http.Request) (int, error) {
	photo, err := parsePhoto(r)
	if err != nil {
		return 500, err
	}
	err = dao.Imp.UpsertPhoto(photo)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func deleteAlbum(r *http.Request) (int, error) {
	theme := getTheme(r)
	err := dao.Imp.ClearAlbumer(theme)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func deletePhoto(r *http.Request) (int, error) {
	id := getPhotoId(r)
	err := dao.Imp.ClearPhoto(id)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

// getAlbum 查询所有theme
func getAlbums(r *http.Request) (*[]album.Album, error) {
	// var group string
	// values := r.URL.Query()
	// group = values.Get("groupName")

	// group, err := getGroup(r)
	// if err != nil {
	// 	return nil, err
	// }

	albums, err := dao.Imp.GetAlbumers()
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// getPhotos 查询某一theme下的所有photo
func getPhotos(r *http.Request) (*[]album.Photo, error) {
	var theme string
	values := r.URL.Query()
	theme = values.Get("theme")

	photos, err := dao.Imp.GetPhotos(theme)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

// getAction 获取action
func getAction(r *http.Request) string {
	// decoder := json.NewDecoder(r.Body)
	// body := make(map[string]interface{})
	// if err := decoder.Decode(&body); err != nil {
	// 	return "", err
	// }
	// defer r.Body.Close()

	// action, ok := body["action"]
	// if !ok {
	// 	return "", fmt.Errorf("缺少 action 参数")
	// }

	// return action.(string), nil

	var action string
	values := r.URL.Query()
	action = values.Get("action")
	return action
}

// 获取url中的theme
func getTheme(r *http.Request) string {
	var theme string
	values := r.URL.Query()
	theme = values.Get("theme")
	return theme
}

// 获取url中的photoId
func getPhotoId(r *http.Request) string {
	var id string
	values := r.URL.Query()
	id = values.Get("photoId")
	return id
}

// parseAlbum 解析Album
func parseAlbum(r *http.Request) (*album.Album, error) {
	decoder := json.NewDecoder(r.Body)
	//body := make(map[string]interface{})
	album := &album.Album{}
	if err := decoder.Decode(album); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return album, nil
}

// parsePhoto 解析Photo body
func parsePhoto(r *http.Request) (*album.Photo, error) {
	decoder := json.NewDecoder(r.Body)
	photo := &album.Photo{}
	if err := decoder.Decode(photo); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return photo, nil
}

// getGroup 获取group
func getGroup(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}
	defer r.Body.Close()

	group, ok := body["groupName"]
	if !ok {
		return "", fmt.Errorf("缺少 group 参数")
	}

	return group.(string), nil
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PhotoHandler photo接口
func PhotoHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {
		action := getAction(r)
		if action == "query" {
			photos, err := getPhotos(r)
			if err != nil {
				res.Code = 404
				res.ErrorMsg = err.Error()
			} else {
				res.Data = photos
			}
		} else if action == "delete" {
			code, err := deletePhoto(r)
			if err != nil {
				res.Code = code
				res.ErrorMsg = err.Error()
			} else {
				res.Data = code
			}
		}
	} else if r.Method == http.MethodPost {
		code, err := modifyPhoto(r)
		if err != nil {
			res.Code = code
			res.ErrorMsg = err.Error()
		} else {
			res.Data = code
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
