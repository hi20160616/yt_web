// template usage example: https://play.golang.org/p/yQzaUEypTe2
package render

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

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
	return v.Description[:600]
}

func Combin() error {
	pattern := "../../../templates/default/tmpl/*.html"
	tmpls, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}
	tmpl := template.Must(tmpls.Clone())
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	return nil
}

// Derive derive values to be displayed.
func (pc *PageCid) Derive(w io.Writer, filename ...string) error {
	overlay := `{{define "title"}}{{.ChannelName}}{{end}}`
	funcs := template.FuncMap{
		"summary": summary,
	}
	temp, err := template.ParseFiles(filename...)
	if err != nil {
		return err
	}
	overlayTmpl, err := template.Must(
		temp.Funcs(funcs).Clone()).
		Parse(overlay)
	if err != nil {
		return err
	}
	if err = overlayTmpl.Execute(w, pc); err != nil {
		return err
	}
	return nil
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
