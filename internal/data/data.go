package data

import (
	"context"
	"strconv"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"google.golang.org/grpc"
)

type Handler struct {
	conn   *grpc.ClientConn
	client pb.YoutubeFetcherClient
}

func (h *Handler) Init() error {
	var err error
	h.conn, err = grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	h.client = pb.NewYoutubeFetcherClient(h.conn)
	return nil
}

func (h *Handler) ChannelName(cid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*3*time.Second)
	defer cancel()
	c, err := h.client.GetChannel(ctx, &pb.Channel{Id: cid})
	if err != nil {
		return "", err
	}
	return c.Name, nil
}

func (h *Handler) Channel(c *pb.Channel) (*pb.Channel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*3*time.Second)
	defer cancel()
	return h.client.GetChannel(ctx, c)
}

func (h *Handler) Channels(cs *pb.Channels) (*pb.Channels, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return h.client.GetChannels(ctx, cs)
}

func (h *Handler) Videos(cid string) (*pb.Videos, error) {
	// only 30 videos caught every time
	ctx, cancel := context.WithTimeout(context.Background(), 30*3*time.Second)
	defer cancel()

	vs, err := h.client.GetVideos(ctx, &pb.Channel{Id: cid})
	if err != nil {
		return nil, err
	}
	// fmt timestamp to RFC3339
	for i := range vs.Videos {
		// nano second dealer
		// t, err := strconv.ParseInt(vs.Videos[i].LastUpdated+"000", 10, 64)
		// vs.Videos[i].LastUpdated = time.Unix(0, t).String()

		// second dealer
		t, err := strconv.ParseInt(vs.Videos[i].LastUpdated[:10], 10, 64)
		if err != nil {
			return nil, err
		}
		vs.Videos[i].LastUpdated = time.Unix(t, 0).String()
	}
	return vs, nil
}

// video assume timeout is 10" for a video
func (h *Handler) Video(vid string) (*pb.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return h.client.GetVideo(ctx, &pb.Video{Id: vid})
}

func (h *Handler) Close() error {
	return h.conn.Close()
}
