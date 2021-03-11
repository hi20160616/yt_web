package main

import (
	"html/template"
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
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		defer h.Close()
		vs := &pb.Videos{}
		vs, err := h.VideosFromTo(vs)
		if err != nil {
			log.Println(err)
		}

		opts := &render.Opts{
			Title: "Home",
			Data:  vs,
			Funcs: template.FuncMap{"summary": render.Summary},
			Tmpls: []string{"layout", "navbar", "index"},
		}
		if err = render.Derive(w, opts); err != nil {
			log.Println(err)
		}
	}
	http.HandleFunc("/", indexHandler)

	channelsHandler := func(w http.ResponseWriter, r *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		defer h.Close()
		cs := &pb.Channels{}
		cs, err := h.Channels(cs)
		if err != nil {
			log.Println(err)
		}

		render.Derive(w, &render.Opts{
			Title: "Channels",
			Data:  cs,
			Tmpls: []string{"layout", "navbar", "channels"}})
	}
	http.HandleFunc("/channels/", channelsHandler)

	videoHandler := func(w http.ResponseWriter, req *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		defer h.Close()
		p := strings.Split(req.URL.Path, "/")
		if p[1] == "vid" {
			res, err := h.Video(p[2])
			if err != nil {
				log.Println(err)
			}
			opts := &render.Opts{
				Title: res.Title,
				Data:  res,
				Tmpls: []string{"layout", "navbar", "vid"},
			}
			if err = render.Derive(w, opts); err != nil {
				log.Println(err)
			}
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
		c, err := h.Channel(&pb.Channel{Id: u[2]})
		if err != nil {
			log.Println(err)
		}
		vs, err := h.Videos(u[2])
		if err != nil {
			log.Println(err)
		}
		p := &render.PageCid{
			ChannelName: c.Name,
			Videos:      vs.Videos,
		}
		funcMap := template.FuncMap{"summary": render.Summary}

		opts := &render.Opts{
			Title: p.ChannelName,
			Data:  p,
			Funcs: funcMap,
			Tmpls: []string{"layout", "navbar", "cid"},
		}
		if err = render.Derive(w, opts); err != nil {
			log.Println(err)
		}
	}
	http.HandleFunc("/cid/", cidHandler)

	searchHandler := func(w http.ResponseWriter, r *http.Request) {
		h := new(data.Handler)
		if err := h.Init(); err != nil {
			log.Println(err)
		}
		defer h.Close()

		u := r.URL.Query().Get("search")
		kws := strings.Split(u, " ")
		vs, err := h.SearchVideos(&pb.Videos{Keywords: kws})
		if err != nil {
			log.Println(err)
		}
		funcMap := template.FuncMap{"summary": render.Summary}

		opts := &render.Opts{
			Title: "Search Result",
			Data:  vs,
			Funcs: funcMap,
			Tmpls: []string{"layout", "navbar", "search"},
		}
		if err = render.Derive(w, opts); err != nil {
			log.Println(err)
		}
	}
	http.HandleFunc("/search/", searchHandler)

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
