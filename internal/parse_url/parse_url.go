// Package parse_url provides a function for parsing and normalization of passed url
package parse_url

import (
	"net/url"
	"strings"

	"github.com/aantoschuk/feed/internal/apperr"
)

// PraseUrl function parses the provided URL string and returns the parsed string.
// Under the hood utilizes net/url to check string for possible errors
// and strip unneccessary data.
//
// It would return an empty string and an error variable if error happened.
func ParseUrl(raw string) (string, error) {
	if strings.TrimSpace(raw) == "" {
		errEmptyString := apperr.NewUserError(
			"passed an empty url",
			"EMPTY_STRING_VALUE",
			1)
		return "", errEmptyString
	}
	u, err := url.Parse(raw)
	if err != nil {
		errParseUrl := apperr.NewInternalError(
			"cannot parse raw url",
			"LIBRARY_ERROR",
			1,
			err)
		return "", errParseUrl
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		errBadScheme := apperr.NewUserError(
			"bad URL scheme. Proper scheme should have http or https",
			"INVALID_VALUE",
			1)
		return "", errBadScheme
	}

	if u.Host == "" {
		errNoHost := apperr.NewUserError(
			"url is missing a host",
			"INVALID_VALUE",
			1)
		return "", errNoHost
	}

	clean := *u
	clean.Fragment = ""
	parsed := clean.String()

	return parsed, nil
}
