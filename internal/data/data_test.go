package data

import (
	"fmt"
	"testing"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
)

func TestVideo(t *testing.T) {
	h := &Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}

	v, err := h.Video("DRZ3tWu-DzY")
	if err != nil {
		t.Error(err)
	}
	if v.Title != "Learn to Pick a Lock with Paperclips" {
		t.Errorf("want title Learn to Pick a Lock with Paperclips, got: %v", v.Title)
	}
}

func TestVideos(t *testing.T) {
	h := &Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}

	vs, err := h.Videos("UCMUnInmOkrWN4gof9KlhNmQ")
	if err != nil {
		t.Error(err)
	}
	for _, v := range vs.Videos {
		fmt.Println(v.Title)
	}
}

func TestGetChannel(t *testing.T) {
	h := &Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}
	cid := "UCMUnInmOkrWN4gof9KlhNmQ"
	channel, err := h.Channel(&pb.Channel{Id: cid})
	if err != nil {
		t.Error(err)
	}
	want := "Mr & Mrs Gao"
	if channel.Name != want {
		t.Errorf("want: %v, got: %v", want, channel.Name)
	}
}

func TestGetChannels(t *testing.T) {
	h := &Handler{}
	if err := h.Init(); err != nil {
		t.Error(err)
	}
	cs, err := h.Channels(&pb.Channels{})
	if err != nil {
		t.Error(err)
	}
	for _, c := range cs.Channels {
		fmt.Println(c)
	}
}
