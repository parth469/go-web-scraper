package scraper

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/parth469/go-web-scraper/internal/model"
	"github.com/parth469/go-web-scraper/utils/config"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func Init() error {
	c := colly.NewCollector()
	url := config.Env.Url

	var fullEvents []model.Event
	var currentEventDate time.Time
	c.OnHTML("tbody", func(body *colly.HTMLElement) {
		var dailyEvent model.Event
		var timeString string

		body.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			parsedDate := ProcessDate(i, tr)
			if !parsedDate.IsZero() {
				currentEventDate = parsedDate
				return
			}

			if !currentEventDate.IsZero() {
				if tr.ChildText("td:nth-child(1)") != "" {
					timeString = tr.ChildText("td:nth-child(1)")
				}

				dailyEvent = ProcessDailyEvent(currentEventDate, tr, timeString)
				fullEvents = append(fullEvents, dailyEvent)
			}
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		return
	})

	c.OnScraped(func(r *colly.Response) {
		if err := helper.SaveToFile(fullEvents); err != nil {
			helper.Log.Error("failed to save events to file", err)
		}
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	return nil
}
