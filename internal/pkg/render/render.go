package render

import (
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"io"
	"os"
)

type Body struct {
	Title  string
	Videos []*pb.Video
}

// Contents get contents from filename
// Copy from https://golang.org/doc/effective_go#defer
// Also this is the best practice of read file effectivily
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}
