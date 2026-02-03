// extractors package helds all specific extractors.
package extractors

import (
	"strings"
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/apperr"
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

	//	get list of elemets with articles
	elements, err := page.Elements(".item-body")
	if err != nil {
		appErr := apperr.NewInternalError("cannot retrieve elemenets from the page", "ELEMENTS_RETRIEVAL_ERROR", 1, err)
		return nil, appErr
	}

	i.Logger.Info("start scrolling")

	for range 2 {
		page.Mouse.Scroll(0, 300., 4)
		page.WaitIdle(2 * time.Second)

		newEls, err := page.Elements(".item-body")
		if err != nil {
			appErr := apperr.NewInternalError("cannot retrieve elemenets from the page", "ELEMENTS_RETRIEVAL_ERROR", 1, err)
			return nil, appErr
		}
		elements = append(elements, newEls...)
	}

	// get base url for the ign, to later get something like  ign/article
	// instead of ign/news/article
	// for some reason they decide to do that way
	clearU := stripUrl(i.URL)
	articles := make([]domain.Article, 30, 30)

	i.Logger.Info("normalizing content")
	// iterate over all elements
	for idx, a := range elements {
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

		link := clearU + *href
		articles[idx] = domain.Article{Url: link, Title: title}

	}

	i.Logger.Info("ign extractor work done")
	return articles, nil
}

// Url satisfies Extractor interface.
// Function getter to get url.
func (i *IGNExtractor) Url() string {
	return i.URL
}
