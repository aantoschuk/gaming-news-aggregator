package extractors

import (
	"fmt"
	"time"

	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

type GamespotExtractor struct {
	URL      string
	WaitTime time.Duration
}

func (g *GamespotExtractor) Extract(page *rod.Page) ([]domain.Article, error) {
	fmt.Println("START: gamestop extractor")

	if err := page.WaitStable(g.WaitTime); err != nil {
		return nil, fmt.Errorf("unexpected error during page loading: %v", err)
	}

	elements, err := page.Elements(".card-item__link")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve elements: %v", err)
	}

	l := len(elements)
	articles := make([]domain.Article, l, l)
	clearU, nil := stripUrlFromSuffix(g.URL, "news")

	for idx, a := range elements {
		href, err := a.Attribute("href")
		if err != nil {
			continue
		}

		h4, err := a.Elements("h4")
		if err != nil {
			continue
		}
		title, err := h4[0].Text()
		if err != nil {
			continue
		}

		link := *clearU + *href
		article := domain.Article{
			Url:   link,
			Title: title,
		}

		articles[idx] = article
	}

	fmt.Println("DONE: gamestop extractor")
	return articles, nil
}

func (g *GamespotExtractor) Url() string {
	return g.URL
}
