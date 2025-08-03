package scraper

import (
	"time"

	"github.com/gocolly/colly"
)

func ProcessDate(i int, tr *colly.HTMLElement) time.Time {
	isTimeSelection := true
	var DateTime time.Time

	tr.ForEach("td", func(i int, td *colly.HTMLElement) {
		if i > 0 {
			isTimeSelection = false
			return
		}
	})

	if isTimeSelection {
		timeString := tr.ChildText("td p")
		layout := "Monday, January 2, 2006"
		DateTime, _ = time.Parse(layout, timeString)
	}

	return DateTime
}
