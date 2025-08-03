package scraper

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/parth469/go-web-scraper/internal/model"
	"github.com/parth469/go-web-scraper/utils/config"
)

func Init() error {
	c := colly.NewCollector()
	url := config.Env.Url

	var eventList []model.Event

	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			isTimeSelection := true
			var DateTime time.Time

			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				if td.Index > 0 {
					isTimeSelection = false
					return
				}
			})

			if isTimeSelection {
				timeString := tr.ChildText("td p")
				layout := "Monday, January 2, 2006"
				DateTime, _ = time.Parse(layout, timeString)
			}

			if !DateTime.IsZero() {
				fmt.Println(DateTime.Format("Monday, January 2, 2006"))
			}
			// event.Time = tr.ChildText("td:nth-child(1)")
			// event.Title = tr.ChildText("td:nth-child(2)")
			// event.EventType = tr.ChildText("td:nth-child(3)")
			// event.Location = tr.ChildText("td:nth-child(4)")
		})
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(eventList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal events to JSON: %w", err)
	}

	err = os.WriteFile("events.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write events.json: %w", err)
	}

	return nil
}
