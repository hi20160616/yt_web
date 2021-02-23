package render

import (
	"io"
	"path"
	"text/template"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

type Page struct {
	Title  string
	Videos []*pb.Video
}

// Derive derive values to be displayed.
func (p *Page) Derive(w io.Writer, filename string) error {
	t := template.Must(template.New(path.Base(filename)).ParseFiles(filename))
	if err := t.Execute(w, p); err != nil {
		return err
	}
	return nil
}
