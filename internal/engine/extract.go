package engine

import (
	"fmt"
	"strings"
	"sync"

	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type extractResult struct {
	articles []domain.Article
	err      error
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

// Extract function prepare browser, page and runs specific Extractors.
// Handles all the logic with outside of the page / Extractor specifics.
//
// returns list of all articles combined from the runned Extractors and an error.
// if error occured, function returns nil and an error variable of type error.
func (e *Engine) Extract() ([]domain.Article, error) {
	browser, err := e.BrowserFactory()
	if err != nil {
		return nil, err
	}
	defer browser.Close()

	defer browser.Close()

	articlesCh := make(chan []domain.Article, len(e.Extractors))
	errCh := make(chan error, len(e.Extractors))
	sem := make(chan struct{}, e.MaxConcurrentJobs)

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
