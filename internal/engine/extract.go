package engine

import (
	"fmt"
	"strings"
	"sync"

	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

type extractResult struct {
	articles []domain.Article
	err      error
}

func workerPool(ex []domain.Extractor, n int, articlesCh chan<- []domain.Article, errCh chan<- error, browser *rod.Browser) {
	var wg sync.WaitGroup
	wg.Add(len(ex))

	ch := make(chan domain.Extractor)

	for i := 0; i < n; i++ {
		go func() {
			for extractor := range ch {
				articles, err := worker(extractor, browser)
				if err != nil {
					errCh <- fmt.Errorf("extractor failed: %v, with error: %w", extractor.Url(), err)
				}
				articlesCh <- articles
				wg.Done()
			}
		}()
	}

	for _, e := range ex {
		ch <- e
	}

	close(ch)

	go func() {
		wg.Wait()
		close(articlesCh)
		close(errCh)
	}()
}

func worker(ex domain.Extractor, browser *rod.Browser) ([]domain.Article, error) {
	url := ex.Url()
	page := browser.MustPage(url)
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
		return nil, fmt.Errorf("failed to start a browser: %v", err)
	}
	defer browser.Close()
	n := len(e.Extractors)

	articlesCh := make(chan []domain.Article)
	errCh := make(chan error, n)
	workerPool(e.Extractors, e.MaxConcurrentJobs, articlesCh, errCh, browser)

	var errors []string

	done := make(chan struct{})
	go func() {
		for err := range errCh {
			errors = append(errors, err.Error())
		}
		close(done)
	}()

	// collect articles
	articles := make([]domain.Article, 0, n*2)
	for a := range articlesCh {
		articles = append(articles, a...)
	}

	<-done
	if len(errors) > 0 {
		// TODO: change to string builder
		return articles, fmt.Errorf("extraction errors: %s", strings.Join(errors, ";\n"))
	}

	return articles, nil
}
