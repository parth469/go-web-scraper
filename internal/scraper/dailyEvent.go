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

	startTime, err := parseFlexibleTime(startTimeStr)
	if err != nil {
		fmt.Println("ERROR")
		return event
	}

	endTime, err := parseFlexibleTime(endTimeStr)
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

func parseFlexibleTime(timeStr string) (time.Time, error) {
	timeStr = strings.ReplaceAll(timeStr, "\u00a0", " ")
	timeStr = strings.ToLower(timeStr)

	formats := []string{
		"3:04 pm", // 11:30 am
		"3:04pm",  // 11:30am
		"3 pm",    // 11 am
		"3pm",     // 11am
		"15:04",   // 09:30 (24-hour format)
		"3:04",    // 9:30 (without am/pm)
	}

	timeStr = strings.ToLower(timeStr)

	// Try each format until one works
	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			// If the format didn't specify am/pm and time is < 12,
			// we need to check if it should be pm
			if !strings.Contains(format, "pm") && !strings.Contains(format, "am") {
				if strings.Contains(timeStr, "pm") && t.Hour() < 12 {
					t = t.Add(12 * time.Hour)
				}
				if strings.Contains(timeStr, "am") && t.Hour() == 12 {
					t = t.Add(-12 * time.Hour)
				}
			}
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unrecognized time format: %s", timeStr)
}
