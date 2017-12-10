package youtube

import (
	"fmt"

	youtube "google.golang.org/api/youtube/v3"
)

// Searcher performs searches on YouTube.
type Searcher interface {
	Search(text string) (SearchResult, error)
}

// SearchResult provides the info a YouTube search result.
type SearchResult interface {
	VideoID() string
}

// Search does a search on YouTube and returs a result or error.
func (c *Client) Search(text string) (SearchResult, error) {
	response, err := c.service.Search.List("snippet").
		Q(text).
		MaxResults(1).
		SafeSearch("none").
		Type("video").
		Do()
	if err != nil {
		return nil, err
	}
	size := len(response.Items)
	if size == 0 {
		return nil, nil
	}
	if size > 1 {
		return nil, fmt.Errorf("unexpected amount of youtube search results: %d", size)
	}

	result := response.Items[0]
	if result.Id == nil {
		return nil, fmt.Errorf("youtube result came without id section")
	}

	return &searchResult{
		response: response,
		result:   result,
	}, nil
}

type searchResult struct {
	response *youtube.SearchListResponse
	result   *youtube.SearchResult
}

var _ SearchResult = (*searchResult)(nil)

// VideoID returns video ID that can be used
// to play the video.
func (s *searchResult) VideoID() string {
	return s.result.Id.VideoId
}
