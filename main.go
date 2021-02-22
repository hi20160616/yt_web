package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hi20160616/yt_web/internal/data"
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
		p := strings.Split(req.URL.Path, "/")
		if p[1] == "cid" {
			res, err := h.Videos(p[2])
			if err != nil {
				log.Println(err)
			}
			for _, video := range res.Videos {
				io.WriteString(w, video.Title)
				io.WriteString(w, "\n")
			}

		}
	}
	http.HandleFunc("/cid/", channelHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
