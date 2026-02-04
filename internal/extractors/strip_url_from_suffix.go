package extractors

import (
	"fmt"
	"strings"
)

// stripUrlFromSuffix removes suffix from the url keeping everything is intact.
func stripUrlFromSuffix(raw, suffix string) (*string, error) {
	if raw == "" || suffix == "" {
		return nil, fmt.Errorf("one or both parameters empty")
	}
	t := "/" + suffix
	if raw[len(raw)-1] == '/' {
		r := strings.TrimSuffix(raw, t +"/")
		return &r, nil
	}
	r := strings.TrimSuffix(raw, t)
	return &r, nil
}
