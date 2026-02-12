package browser

import (
	"time"

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

// InfiniteScroll function allows to scroll and retrieve html elements from the page.
// Controls how many times to scroll, wait per scroll and what dom element to select.
// return browser elements or error so that could futher do something with it.
func InfiniteScroll(page *rod.Page, scrolls int, scrollPixels float64, waitPerScroll time.Duration, selector string) (rod.Elements, error) {
	for range scrolls {
		page.Mouse.Scroll(0, scrollPixels, 100)
		page.WaitIdle(waitPerScroll)
	}

	elements, err := page.Elements(selector)
	if err != nil {
			return nil, err
		}
	return elements, nil
}
