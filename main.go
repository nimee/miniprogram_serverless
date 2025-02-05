package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	// http.HandleFunc("/", service.IndexHandler)
	// http.HandleFunc("/api/count", service.CounterHandler)

	http.HandleFunc("/api/album", service.AlbumHandler)
	http.HandleFunc("/api/photo", service.PhotoHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
