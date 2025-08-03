package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/parth469/go-web-scraper/internal/model"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func ProcessDailyEvent(date time.Time, tr *colly.HTMLElement, timeStr string) model.Event {
	var event model.Event
	event.Title = helper.CleanString(tr.ChildText("td:nth-child(2)"))
	event.EventType = helper.CleanString(tr.ChildText("td:nth-child(3)"))
	event.Location = helper.CleanString(tr.ChildText("td:nth-child(4)"))

	timeParts := strings.Split(timeStr, "-")
	startTimeStr := strings.TrimSpace(timeParts[0])
	endTimeStr := strings.TrimSpace(timeParts[1])

	startTime, err := helper.ParseFlexibleTime(startTimeStr)
	if err != nil {
		fmt.Println("ERROR")
		return event
	}

	endTime, err := helper.ParseFlexibleTime(endTimeStr)
	if err != nil {
		fmt.Println("ERROR")
		return event
	}

	event.StartTime = time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		startTime.Hour(),
		startTime.Minute(),
		0, 0, date.Location(),
	)

	event.EndTime = time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		endTime.Hour(),
		endTime.Minute(),
		0, 0, date.Location(),
	)

	return event

}
