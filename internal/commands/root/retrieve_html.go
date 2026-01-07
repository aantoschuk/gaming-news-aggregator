package root

import (
	"strings"

	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// RetrieveHTML scraps html content from the page.
// Accepts url as a string, return an error.
func RetrieveHTML(url string) ([]article, error) {

	s := strings.Split(url, "/")
	cleared := strings.Join(s[:len(s)-1], "/")
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	opts := proto.TargetCreateTarget{
		URL: url,
	}
	page, err := browser.Page(opts)
	var articles []article
	if err != nil {
		if strings.Contains(err.Error(), "ERR_INTERNET_DISCONNECTED") {
			return articles, apperr.ErrNoInternetConnection
		}
		appErr := apperr.NewInternalError("cannot create a tab", "LIBRARY_ERROR", 1, err)
		return articles, appErr
	}

	page = page.MustWaitStable()
	if err != nil {
		appErr := apperr.NewInternalError("error while waiting a stable page", "LIBRARY_ERROR", 1, err)
		return articles, appErr
	}

	elements, err := page.Elements(".item-body")
	if err != nil {
		appErr := apperr.NewInternalError("cannot get the element", "LIBRARY_ERROR", 1, err)
		return articles, appErr
	}

	t := article{}

	for _, a := range elements {

		href, err := a.Attribute("href")
		if err != nil || href == nil {
			continue
		}
		div, err := a.Elements("div")
		if err != nil {
			continue
		}

		span, _ := div[1].Element("span")
		text, _ := span.Text()

		link := *href
		t.url = cleared + link
		t.title = text
		articles = append(articles, t)
	}

	return articles, nil
}

// TODO: refactor. Handle all errors. Make more modular, 
