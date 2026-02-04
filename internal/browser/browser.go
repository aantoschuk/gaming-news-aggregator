package browser

import (
	"time"

	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// InitBrowser creates a new browser process.
// Headless or not depending on the isDebug parameter.
func InitBrowser(delay time.Duration, isDebug bool) *rod.Browser {
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

func InfiniteScroll(page *rod.Page, scrolls int, scrollPixels float64, waitPerScroll time.Duration, selector string) (rod.Elements, error) {
	for range scrolls {
		page.Mouse.Scroll(0, scrollPixels, 100)
		page.WaitIdle(waitPerScroll)
	}

	elements, err := page.Elements(selector)
	if err != nil {
		return nil, apperr.NewInternalError(
			"cannot retrieve elements from the page",
			"ELEMENTS_RETRIEVAL_ERROR",
			1,
			err,
		)
	}

	return elements, nil
}
