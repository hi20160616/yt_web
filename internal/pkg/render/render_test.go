package render

import (
	"html/template"
	"os"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_web/internal/data"
)

// func TestDeriveCid(t *testing.T) {
//         h := &data.Handler{}
//         if err := h.Init(); err != nil {
//                 t.Error(err)
//         }
//         cid := "UCCtTgzGzQSWVzCG0xR7U-MQ"
//         v, err := h.Videos(cid)
//         if err != nil {
//                 t.Error(err)
//         }
//
//         cname, err := h.ChannelName(cid)
//         if err != nil {
//                 t.Error(err)
//         }
//         p := &PageCid{ChannelName: cname}
//         p.Videos = v.Videos
//         p.ChannelName = cname
//         if err = p.Derive(os.Stdout, "../../../templates/default/cid.html"); err != nil {
//                 t.Error(err)
//         }
// }

func TestDerive(t *testing.T) {
	h := new(data.Handler)
	if err := h.Init(); err != nil {
		t.Error(err)
	}

	cid := "UCCtTgzGzQSWVzCG0xR7U-MQ"
	c := &pb.Channel{Id: cid}
	c, err := h.Channel(c)
	if err != nil {
		t.Error(err)
	}
	v, err := h.Videos(cid)
	if err != nil {
		t.Error(err)
	}
	p := &PageCid{
		ChannelName: c.Name,
		Videos:      v.Videos,
	}
	summary := func(v *pb.Video) string {
		return v.Description[:600]
	}
	funcMap := template.FuncMap{"summary": summary}

	opts := &Opts{
		Title: p.ChannelName, // TODO: rm p
		Data:  p,
		Funcs: funcMap,
		Tmpls: []string{"layout", "navbar", "cid"},
	}

	if err = Derive(os.Stdout, opts); err != nil {
		t.Error(err)
	}
}
