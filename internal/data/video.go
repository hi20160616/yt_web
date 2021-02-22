package data

import (
	"context"
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

func (h *Handler) videos(cid string) (*pb.Videos, error) {
	// only 30 videos caught every time
	ctx, cancel := context.WithTimeout(context.Background(), 30*3*time.Second)
	defer cancel()

	return h.client.GetVideos(ctx, &pb.Channel{Cid: cid})
}

// video assume timeout is 10" for a video
func (h *Handler) video(vid string) (*pb.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return h.client.GetVideo(ctx, &pb.Video{Vid: vid})
}

func (h *Handler) Close() error {
	return h.conn.Close()
}
