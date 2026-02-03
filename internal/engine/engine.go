// engine package is hold a structure which collects different extractors
// and run them, to retrieve information.
package engine

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
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
	Debug             bool
	MaxConcurrentJobs int
	// controls how much slow motion would be
	//TODO: Not sure SlowMotion is needed at all, as there is no browser acttions.
	// To slow retrieving content, better to use sleep in the Extractor itself.
	SlowMotion time.Duration
}

type extractResult struct {
	articles []domain.Article
	err      error
}

// Extract function prepare browser, page and runs specific Extractors.
// Handles all the logic with outside of the page / Extractor specifics.
//
// returns list of all articles combined from the runned Extractors and an error.
// if error occured, function returns nil and an error variable of type error.
func (e *Engine) Extract() ([]domain.Article, error) {
	browser := initBrowser(e.SlowMotion, e.Debug)
	defer browser.Close()

	concurrency := max(e.MaxConcurrentJobs, 1)
	concurrency = min(concurrency, len(e.Extractors))

	articlesCh := make(chan []domain.Article, len(e.Extractors))
	errCh := make(chan error, len(e.Extractors))
	sem := make(chan struct{}, concurrency)

	var wg sync.WaitGroup

	for _, ex := range e.Extractors {
		wg.Add(1)
		ex := ex

		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			extracted, err := extractContent(ex, browser)
			if err != nil {
				errCh <- fmt.Errorf("extractor %v failed: %w", ex.Url(), err)
				return
			}
			articlesCh <- extracted
		}()
	}

	wg.Wait()
	close(articlesCh)
	close(errCh)

	// collect articles
	var articles []domain.Article
	for a := range articlesCh {
		articles = append(articles, a...)
	}

	// error handling
	var errMsgs []string
	for err := range errCh {
		errMsgs = append(errMsgs, err.Error())
	}
	if len(errMsgs) > 0 {
		return articles, fmt.Errorf("extraction errors: %s", strings.Join(errMsgs, "; "))
	}

	return articles, nil
}

func extractContent(ex domain.Extractor, browser *rod.Browser) ([]domain.Article, error) {
	url := ex.Url()
	page, err := browser.Page(proto.TargetCreateTarget{URL: url})

	if err != nil {
		return nil, err
	}

	defer page.MustClose()

	extracted, err := ex.Extract(page)
	if err != nil {
		return nil, err
	}

	return extracted, nil
}

func initBrowser(delay time.Duration, isDebug bool) *rod.Browser {
	var browser *rod.Browser
	if isDebug {
		l := launcher.New().Headless(false).Devtools(true)
		defer l.Cleanup()
		url := l.MustLaunch()
		browser = rod.New().ControlURL(url).Trace(true).SlowMotion(delay).MustConnect()
	} else {
		browser = rod.New().MustConnect()
	}
	return browser
}
