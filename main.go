package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_web/internal/data"
	"github.com/hi20160616/yt_web/internal/pkg/render"
)

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

	cidHandler := func(w http.ResponseWriter, req *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		defer h.Close()
		u := strings.Split(req.URL.Path, "/")
		if u[1] == "cid" {
			c, err := h.Channel(&pb.Channel{Id: u[2]})
			if err != nil {
				log.Println(err)
			}
			res, err := h.Videos(u[2])
			if err != nil {
				log.Println(err)
			}
			p := &render.Page{
				ChannelName: c.Name,
				Videos:      res.Videos,
			}
			if err = p.Derive(w, "./templates/default/cid.html"); err != nil {
				log.Println(err)
			}
		}
	}
	http.HandleFunc("/cid/", cidHandler)

	staticHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		u := strings.Split(r.URL.Path, "/")
		if u[1] == "static" {
			if u[2] == "css" {
				raw, err := ioutil.ReadFile("./templates/default/" + u[3])
				if err != nil {
					log.Println(err)
				}
				io.WriteString(w, string(raw))
			}

		}
		io.Copy(w, r.Body)
	}
	http.HandleFunc("/static/css/", staticHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
