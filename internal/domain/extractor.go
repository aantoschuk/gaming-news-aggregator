// package domain defines models or types used by the app.
package domain

import (
	"github.com/go-rod/rod"
)

// Extractor defines the core behavior for all specific extractors.
//
// Any Extractor must provide:
// - Extract: logic to extract articles from a given page
// - Url: getter for the URL to the page associated with the extractor
type Extractor interface {
	Extract(page *rod.Page) ([]Article, error)
	Url() string
}
