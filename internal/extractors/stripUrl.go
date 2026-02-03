package extractors

import "strings"

// remove /news from the url
func stripUrl(raw string) string {
	s := strings.Split(raw, "/")
	return strings.Join(s[:len(s)-1], "/")
}
