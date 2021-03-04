// template usage example: https://play.golang.org/p/yQzaUEypTe2
package render

import (
	"fmt"
	"html/template"
	"io"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

type PageCid struct {
	ChannelName string
	Videos      []*pb.Video
}

type templateFile struct {
	name     string
	contents string
}

type Opts struct {
	Data  interface{}
	Title string
	Funcs template.FuncMap
	Tmpls []string
}

func summary(v *pb.Video) string {
	dRune := []rune(v.Description)
	if len(dRune) <= 300 {
		return v.Description
	}
	return string(dRune[:300])
}

func Derive(w io.Writer, opts *Opts) error {
	overlayTitle := `{{define "title"}}` + opts.Title + `{{end}}`
	var tmpls []string
	for _, tmpl := range opts.Tmpls {
		tmpls = append(tmpls, fmt.Sprintf("./templates/default/%s.html", tmpl))
		// just for test
		// tmpls = append(tmpls, fmt.Sprintf("../../../templates/default/%s.html", tmpl))
	}
	t := template.New("")
	if opts.Funcs != nil {
		t = t.Funcs(opts.Funcs)
	}
	t = template.Must(t.ParseFiles(tmpls...))
	t, err := t.Parse(overlayTitle)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, "layout", opts.Data)
}
