// extractors package helds all specific extractors.
package extractors

import (
	"strings"
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

// IGNExtractor defines specific Extractor
type IGNExtractor struct {
	URL      string
	WaitTime time.Duration
	Logger   *app_logger.AppLogger
}

// Extract function is implementation of Extract interface.
// It accepts page, extract all articles from the page
// and return them back to the engine.
func (i *IGNExtractor) Extract(page *rod.Page) ([]domain.Article, error) {
	i.Logger.Info("starting ign extractor")

	// wait until page loads
	if err := page.WaitStable(i.WaitTime); err != nil {
		return nil, err
	}

	//	get list of elemets with articles
	elements, err := page.Elements(".item-body")
	if err != nil {
		return nil, err
	}

	// get base url for the ign, to later get something like  ign/article
	// instead of ign/news/article
	// for some reason they decide to do that way
	clearU := stripUrl(i.URL)
	articles := []domain.Article{}

	// iterate over all elements
	for _, a := range elements {
		// get link to the full article
		href, err := a.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		// retrieve nested components
		divs, err := a.Elements("div")
		if err != nil || len(divs) < 2 {
			continue
		}

		span, err := divs[1].Element("span")
		if err != nil {
			continue
		}
		text, err := span.Text()
		if err != nil {
			continue
		}

		link := clearU + *href
		articles = append(articles, domain.Article{Url: link, Title: text})
	}

	i.Logger.Info("ign extractor work done")
	return articles, nil
}

// Url satisfies Extractor interface.
// Function getter to get url.
func (i *IGNExtractor) Url() string {
	return i.URL
}

// remove /news from the url
func stripUrl(raw string) string {
	s := strings.Split(raw, "/")
	return strings.Join(s[:len(s)-1], "/")
}
