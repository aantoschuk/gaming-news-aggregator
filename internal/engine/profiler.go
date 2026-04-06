package engine

import (
	"fmt"
	"sync"
	"time"
)

type Profiler struct {
	startTimes map[string]time.Time
	mu         *sync.Mutex
}

func NewProfiler() *Profiler {
	return &Profiler{
		startTimes: make(map[string]time.Time),
		mu:         &sync.Mutex{},
	}
}

func (p *Profiler) Start(extractorName string) {
	fmt.Printf("Starting: %s\n", extractorName)
	p.mu.Lock()
	p.startTimes[extractorName] = time.Now()
	p.mu.Unlock()
}

func (p *Profiler) End(extractorName string) {
	p.mu.Lock()
	start, ok := p.startTimes[extractorName]
	if ok {
		delete(p.startTimes, extractorName)
	}
	p.mu.Unlock()

	if !ok {
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Done: %s (%s)\n", extractorName, elapsed)
}
