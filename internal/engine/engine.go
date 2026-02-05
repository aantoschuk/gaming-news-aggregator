// engine package is hold a structure which collects different extractors
// and run them, to retrieve information.
package engine

import (
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/browser"
	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

// Engine defines an object which collects Extractors and runs them.
type Engine struct {
	// slice of the all Extractors which would be runned.
	Extractors []domain.Extractor
	// App logger to display messages.
	Logger *app_logger.AppLogger
	// Controls how many extractors will work at the same time. Default value is 1.
	// If value is bigger then 1, it means that the multiple extractors will run.
	// Value cannot be bigger then the amount of passed extractors,
	// otherwise it will be set to that amount.
	MaxConcurrentJobs int
	// controls how much slow motion would be
	//TODO: Not sure SlowMotion is needed at all, as there is no browser acttions.
	// To slow retrieving content, better to use sleep in the Extractor itself.
	SlowMotion time.Duration

	// Accepts a created row browser which will later open pages for every passed extractor
	BrowserFactory func() (*rod.Browser, error)
}

type CreateEngineParams struct {
	Logger            *app_logger.AppLogger
	Extractors        []domain.Extractor
	MaxConcurrentJobs int
	SlowMotion        time.Duration
}

// CreateEngine function allocates
func CreateEngine(params CreateEngineParams) *Engine {
	if len(params.Extractors) == 0 {
		panic("engine requires at least one extractor")
	}

	concurrency := max(params.MaxConcurrentJobs, 1)
	concurrency = min(concurrency, len(params.Extractors))

	factory := func() (*rod.Browser, error) {
		return browser.InitBrowser(0, false), nil
	}

	en := Engine{
		Extractors:        params.Extractors,
		Logger:            params.Logger,
		MaxConcurrentJobs: concurrency,
		BrowserFactory:    factory,
		SlowMotion:        params.SlowMotion,
	}

	return &en
}
