// template usage example: https://play.golang.org/p/yQzaUEypTe2
package render

import (
	"html/template"
	"io"
	"path"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

type Page struct {
	ChannelName string
	Videos      []*pb.Video
}

func summary(v *pb.Video) string {
	return v.Description[:600]
}

// Derive derive values to be displayed.
func (p *Page) Derive(w io.Writer, filename string) error {
	funcs := template.FuncMap{
		"summary": summary,
	}
	report := template.Must(template.New(path.Base(filename)).
		Funcs(funcs).
		ParseFiles(filename))
	if err := report.Execute(w, p); err != nil {
		return err
	}
	return nil
}
