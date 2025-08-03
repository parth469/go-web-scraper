package scraper

import (
	"github.com/gocolly/colly"
	"github.com/parth469/go-web-scraper/internal/model"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func ProcessPoster(link string, event model.Event, subCollector *colly.Collector) {
	var PosterList []model.Poster

	subCollector.OnHTML("tbody", func(body *colly.HTMLElement) {
		body.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			currentPoster := model.Poster{
				StartTime: event.StartTime,
				EndTime:   event.EndTime,
				PosterId:  helper.CleanString(tr.ChildText("td:nth-child(1) p")),
				Title:     helper.CleanString(tr.ChildText("td:nth-child(2) p")),
				Presenter: helper.CleanString(tr.ChildText("td:nth-child(3) p")),
				TopicArea: helper.CleanString(tr.ChildText("td:nth-child(4) p")),
				Abstract:  "",
			}

			PosterList = append(PosterList, currentPoster)
		})
	})

	subCollector.OnError(func(r *colly.Response, e error) {
		return
	})

	subCollector.OnScraped(func(r *colly.Response) {
		if err := helper.SavePoster(PosterList); err != nil {
			helper.Log.Error("failed to save events to file", err)
		}
	})

	err := subCollector.Visit(link)

	if err != nil {
		helper.Log.Error("failed to visit poster link", err)
	}
}
