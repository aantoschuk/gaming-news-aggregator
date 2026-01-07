package root

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// RetrieveHTML scraps html content from the page.
// Accepts url as a string, return an error.
func RetrieveHTML(browser *rod.Browser, url string) (string, error) {

	page, err := loadPage(browser, url)
	if err != nil {
		return "", err
	}

	html, err := page.HTML()
	if err != nil {
		return "", err
	}

	return html, nil
}

// TODO: refactor. Handle all errors. Make more modular,

func loadPage(browser *rod.Browser, url string) (*rod.Page, error) {
	page, err := browser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return nil, err
	}

	if err := page.WaitStable(1 * time.Second); err != nil {
		return nil, err
	}

	return page, nil
}

/*
func getArticles(url string, elements rod.Elements) []article {
	t := article{}
	var articles []article

	cleared := stripUrl(url)
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
	return articles
}
*/
