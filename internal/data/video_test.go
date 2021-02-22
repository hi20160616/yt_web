package data

import "testing"

func TestVideo(t *testing.T) {
	h := &Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}

	v, err := h.video("DRZ3tWu-DzY")
	if err != nil {
		t.Error(err)
	}
	if v.Title != "Learn to Pick a Lock with Paperclips" {
		t.Errorf("want title Learn to Pick a Lock with Paperclips, got: %v", v.Title)
	}
}
