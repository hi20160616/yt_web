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
	v, err := h.Videos("UCCtTgzGzQSWVzCG0xR7U-MQ")
	if err != nil {
		t.Error(err)
	}

	p := &Page{Title: "Youtube video index"}
	p.Videos = v.Videos
	if err = p.Derive(os.Stdout, "../../../templates/default/index.html"); err != nil {
		t.Error(err)
	}
}
