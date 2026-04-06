// engine package is hold a structure which collects different extractors
// and run them, to retrieve information.
package engine

import (
	"time"

	"github.com/aantoschuk/feed/internal/browser"
	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

// Engine defines an object which collects Extractors and runs them.
type Engine struct {
	Extractors        []domain.Extractor
	MaxConcurrentJobs int
	SlowMotion        time.Duration
	Profiler          bool

	// Accepts a created row browser which will later open pages for every passed extractor
	BrowserFactory func() (*rod.Browser, error)
}

// CreateEngine function allocates
func CreateEngine(params CreateEngineParams) *Engine {
	if len(params.Extractors) == 0 {
		panic("engine requires at least one extractor")
	}

	n := len(params.Extractors)
	concurrency := params.MaxConcurrentJobs
	if concurrency == 0 {
		concurrency = n
	} else if params.MaxConcurrentJobs != 1 && params.MaxConcurrentJobs > n {
		concurrency = n
	}

	factory := func() (*rod.Browser, error) {
		return browser.InitBrowser(0, params.Debug), nil
	}

	en := Engine{
		Extractors:        params.Extractors,
		MaxConcurrentJobs: concurrency,
		BrowserFactory:    factory,
		SlowMotion:        params.Browser.SlowMotion,
		Profiler:          params.Profiler.IsEnabled,
	}

	return &en
}
