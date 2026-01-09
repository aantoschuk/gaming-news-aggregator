// engine package is hold a structure which collects different extractors
// and run them, to retrieve information.
package engine

import (
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// Engine defines an object which collects Extractors and runs them.
type Engine struct {
	// slice of the all Extractors which would be runned.
	Extractors []domain.Extractor
	// App logger to display messages.
	Logger *app_logger.AppLogger
	// Enables debug mode which removes headless browser.
	Debug bool
	// controls how much slow motion would be
	//TODO: Not sure SlowMotion is needed at all, as there is no browser acttions.
	// To slow retrieving content, better to use sleep in the Extractor itself.
	SlowMotion time.Duration
}

// Extract function prepare browser, page and runs specific Extractors.
// Handles all the logic with outside of the page / Extractor specifics.
//
// returns list of all articles combined from the runned Extractors and an error.
// if error occured, function returns nil and an error variable of type error.
func (e *Engine) Extract() ([]domain.Article, error) {
	// start a new browser process
	var browser *rod.Browser

	if e.Debug {
		l := launcher.New().Headless(false).Devtools(true)

		defer l.Cleanup()
		url := l.MustLaunch()
		browser = rod.New().ControlURL(url).Trace(true).SlowMotion(e.SlowMotion).MustConnect()

	} else {
		browser = rod.New().MustConnect()
	}

	defer browser.Close()

	articles := []domain.Article{}

	// run all Extractors
	for _, ex := range e.Extractors {
		url := ex.Url()

		// prepare page in the browser
		page, err := browser.Page(proto.TargetCreateTarget{URL: url})
		if err != nil {
			appErr := apperr.NewInternalError("something happened during openning the page", "OPEN_TAB_ERROR", 1, err)
			return nil, appErr
		}
		defer page.MustClose()

		// get articles from the page
		extracted, err := ex.Extract(page)
		if err != nil {
			return nil, err

		}

		articles = append(articles, extracted...)
	}
	return articles, nil
}
