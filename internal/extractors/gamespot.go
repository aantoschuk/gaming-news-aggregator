package extractors

import (
	"time"

	"github.com/aantoschuk/feed/internal/app_logger"
	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/aantoschuk/feed/internal/domain"
	"github.com/go-rod/rod"
)

type GamespotExtractor struct {
	URL      string
	WaitTime time.Duration
	Logger   *app_logger.AppLogger
}

func (g *GamespotExtractor) Extract(page *rod.Page) ([]domain.Article, error) {
	g.Logger.Info("gamestop extractor started ")
	if err := page.WaitStable(g.WaitTime); err != nil {
		appErr := apperr.NewInternalError("something happened during wait time for gamespot page loading", "PAGE_STABLE_ERROR", 1, err)
		return nil, appErr
	}

	elements, err := page.Elements(".card-item__link")
	if err != nil {
		appErr := apperr.NewInternalError("cannot retrieve elements from the gamespot page", "ELEMENT_RETRIEVAL_ERROR", 1, err)
		return nil, appErr
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

	g.Logger.Info("gamespot extracting done")
	return articles, nil
}

func (g *GamespotExtractor) Url() string {
	return g.URL
}
