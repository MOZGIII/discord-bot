package youtube

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

// Client represents a YouTube client.
type Client struct {
	client  *http.Client
	service *youtube.Service
}

var _ Searcher = (*Client)(nil)

// New builds new YouTube client.
func New(key string) (*Client, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}
	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}
	return &Client{
		client:  client,
		service: service,
	}, nil
}
