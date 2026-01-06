// Package parse_url provides a function for parsing and normalization of passed url
package parse_url

import (
	"errors"
	"net/url"
	"strings"
)

// PraseUrl function parses the provided URL string and returns the parsed string.
// Under the hood utilizes net/url to check string for possible errors
// and strip unneccessary data.
//
// It would return an empty string and an error variable if error happened.
func ParseUrl(raw string) (string, error) {
	if strings.TrimSpace(raw)== "" {
		errEmptyString := errors.New("passed an empty url")
		return "", errEmptyString
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		errBadScheme := errors.New("bad URL scheme. Proper scheme should have 'http' or 'https'")
		return "", errBadScheme
	}

	if u.Host == "" {
		errNoHost := errors.New("URL is missing a host")
		return "", errNoHost
	}

	clean := *u
	clean.Fragment = ""
	parsed := clean.String()

	return parsed, nil
}
