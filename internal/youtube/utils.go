package youtube

import (
	"fmt"
	"net/url"
)

// VideoURL makes a YouTube URL for a specified Video ID.
func VideoURL(id string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", url.QueryEscape(id))
}
