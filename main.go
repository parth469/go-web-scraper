package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	if _, err := parseFlexibleTime("11:30 am"); err != nil {
		fmt.Println(err)
	}
}

func parseFlexibleTime(timeStr string) (time.Time, error) {
	timeStr = strings.ToLower(timeStr)
	formats := []string{
		"3:04 pm", "3:04pm", "3 pm", "3pm", // PM formats
		"3:04 am", "3:04am", "3 am", "3am", // AM formats
		"15:04", "3:04", // 24-hour & simple formats
	}

	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			// If format didn't have am/pm but original string did, adjust
			if !strings.Contains(format, "m") {
				if strings.Contains(timeStr, "pm") && t.Hour() < 12 {
					t = t.Add(12 * time.Hour)
				} else if strings.Contains(timeStr, "am") && t.Hour() == 12 {
					t = t.Add(-12 * time.Hour)
				}
			}
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized time format: %s", timeStr)
}
