// extractors package helds all specific extractors.
package extractors

import (
	"strings"
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/aantoschuk/feed/internal/browser"
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
		appErr := apperr.NewInternalError("something happened during wait for page load.", "PAGE_STABLE_ERROR", 1, err)
		return nil, appErr
	}

	elements, err := browser.InfiniteScroll(page, 2, 1000, 1*time.Second, ".item-body")
	if err != nil {
		return nil, err
	}

	// get base url for the ign, to later get something like  ign/article
	// instead of ign/news/article
	// for some reason they decide to do that way
	clearU, err := stripUrlFromSuffix(i.URL, "news")
	articles := []domain.Article{}

	i.Logger.Info("normalizing content")
	// iterate over all elements
	for _, a := range elements {
		// get link to the full article
		href, err := a.Attribute("href")
		if err != nil || href == nil {
			i.Logger.Debug("retrieve attribute error: " + err.Error())
			continue
		}

		// retrieve nested components
		spanStack, err := a.Elements("span")

		if err != nil {
			i.Logger.Debug("retrieve attribute error: " + err.Error())
			continue
		}

		span := spanStack[1]

		text, err := span.Text()
		title := strings.Split(text, "\n")[0]

		if err != nil {
			i.Logger.Debug("cannot get post title: " + err.Error())
			continue
		}

		link := *clearU + *href
		articles = append(articles, domain.Article{Url: link, Title: title})

	}

	i.Logger.Info("ign extractor work done")
	return articles, nil
}

// Url satisfies Extractor interface.
// Function getter to get url.
func (i *IGNExtractor) Url() string {
	return i.URL
}
