package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hi20160616/yt_web/internal/data"
	"github.com/hi20160616/yt_web/internal/pkg/render"
)

// template usage example: https://play.golang.org/p/yQzaUEypTe2

func main() {
	videoHandler := func(w http.ResponseWriter, req *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		p := strings.Split(req.URL.Path, "/")
		if p[1] == "vid" {
			res, err := h.Video(p[2])
			if err != nil {
				log.Println(err)
			}
			io.WriteString(w, res.Title)
			io.WriteString(w, "\n")
			io.WriteString(w, res.Description)
		}
	}
	http.HandleFunc("/vid/", videoHandler)

	channelHandler := func(w http.ResponseWriter, req *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		u := strings.Split(req.URL.Path, "/")
		if u[1] == "cid" {
			res, err := h.Videos(u[2])
			if err != nil {
				log.Println(err)
			}
			p := &render.Page{
				Title:  "Channel id: " + u[1],
				Videos: res.Videos,
			}
			if err = p.Derive(w, "./templates/default/index.html"); err != nil {
				log.Println(err)
			}
		}
	}
	http.HandleFunc("/cid/", channelHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
