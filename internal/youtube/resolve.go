package youtube

import (
	"fmt"
	"net/url"
)

// Resolve parses input and converts it to a YouTube
// video URL.
func (c *Client) Resolve(input string) (string, error) {
	// First try to parse a video URL. If succeeded,
	// return Video ID from the URL.
	if videoURL, err := parseYouTubeVideoURL(input); err == nil {
		return videoURL.VideoID()
	}
	// It's not a URL.
	// Try to search YouTube.
	result, err := c.Search(input)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", fmt.Errorf("no videos found for %q", input)
	}
	return result.VideoID(), nil
}

func parseYouTubeVideoURL(input string) (*youTubeVideoURL, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}
	if u.Hostname() != "youtube.com" {
		return nil, fmt.Errorf("wrong URL host name: %s", u.Hostname())
	}
	return &youTubeVideoURL{url: u}, nil
}

type youTubeVideoURL struct {
	url *url.URL
}

func (u *youTubeVideoURL) VideoID() (string, error) {
	values := u.url.Query()
	v, ok := values["v"]
	if !ok || len(v) == 0 {
		return "", fmt.Errorf("no v param found in URL")
	}
	return v[0], nil
}
