package youtube

import (
	"os"
)

// DefaultClient is a deafult YouTube client.
var DefaultClient *Client

func init() {
	key := os.Getenv("YOUTUBE_KEY")
	client, err := New(key)
	if err != nil {
		panic(err)
	}
	DefaultClient = client
}
