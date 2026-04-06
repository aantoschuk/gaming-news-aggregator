package engine

import (
	"github.com/aantoschuk/feed/internal/domain"
	"time"
)

type CreateEngineParams struct {
	// Slice of the all Extractors which would be runned.
	Extractors []domain.Extractor

	//Responsible for how many extractors will work at the same time. Default value is 1.
	// If value is bigger then 1, it means that the multiple extractors will run.
	// Value cannot be bigger then the amount of passed extractors,
	// otherwise it will be set to that amount.
	MaxConcurrentJobs int

	// Holds runtime settings for the browser used by the Engine.
	Browser Browser

	// Provides optional instrumentation for extractors.
	Profiler ProfilerOptions

	// Launhc browser in non-headless mode.
	Debug bool
}

type Browser struct {
	// Controls how much slow motion would be. Only for the browser actions.
	// To slow retrieving content, better to use sleep in the Extractor itself.
	SlowMotion time.Duration
}

type ProfilerOptions struct {
	// Turns on the profiler for extractors.
	// When true, start / end messages and timing measurements are logged.
	IsEnabled bool
}
