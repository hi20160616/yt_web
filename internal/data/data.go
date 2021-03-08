package data

import (
	"context"
	"sort"
	"strconv"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"google.golang.org/grpc"
)

type Handler struct {
	conn   *grpc.ClientConn
	client pb.YoutubeFetcherClient
}

type ByLastUpdated []*pb.Video

func (b ByLastUpdated) Len() int {
	return len(b)
}

func (b ByLastUpdated) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByLastUpdated) Less(i, j int) bool {
	bi, _ := strconv.ParseInt(b[i].LastUpdated, 10, 64)
	bj, _ := strconv.ParseInt(b[j].LastUpdated, 10, 64)
	return bi < bj
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
	cs, err := h.client.GetChannels(ctx, cs)
	if err != nil {
		return nil, err
	}
	for i := range cs.Channels {
		t, err := strconv.ParseInt(cs.Channels[i].LastUpdated[:10], 10, 64)
		if err != nil {
			return nil, err
		}
		cs.Channels[i].LastUpdated = time.Unix(t, 0).Format("01-02 15:04")
	}
	return cs, nil
}

func (h *Handler) Videos(cid string) (*pb.Videos, error) {
	// only 30 videos caught every time
	ctx, cancel := context.WithTimeout(context.Background(), 30*3*time.Second)
	defer cancel()

	vs, err := h.client.GetVideos(ctx, &pb.Channel{Id: cid})
	if err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(ByLastUpdated(vs.Videos)))
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
		vs.Videos[i].LastUpdated = time.Unix(t, 0).Format("01-02 15:04")
	}
	return vs, nil
}

// video assume timeout is 10" for a video
func (h *Handler) Video(vid string) (*pb.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	v, err := h.client.GetVideo(ctx, &pb.Video{Id: vid})
	t, err := strconv.ParseInt(v.LastUpdated[:10], 10, 64)
	if err != nil {
		return nil, err
	}
	v.LastUpdated = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	return v, nil
}

func (h *Handler) Close() error {
	return h.conn.Close()
}
