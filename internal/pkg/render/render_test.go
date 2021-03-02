package render

import (
	"os"
	"testing"

	"github.com/hi20160616/yt_web/internal/data"
)

func TestParse(t *testing.T) {
	h := &data.Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}
	cid := "UCCtTgzGzQSWVzCG0xR7U-MQ"
	v, err := h.Videos(cid)
	if err != nil {
		t.Error(err)
	}

	cname, err := h.ChannelName(cid)
	if err != nil {
		t.Error(err)
	}
	p := &Page{ChannelName: cname}
	p.Videos = v.Videos
	p.ChannelName = cname
	if err = p.Derive(os.Stdout, "../../../templates/default/index.html"); err != nil {
		t.Error(err)
	}
}
