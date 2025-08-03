package helper

import (
	"fmt"
	"strings"
	"time"
)

func CleanString(input string) string {
	// Define replacements in a map for better maintainability
	replacements := map[string]string{
		"\n":     ". ",  // newline
		"\t":     "",    // tab
		"\u0026": "and", // ampersand
		"\u2011": "-",   // non-breaking hyphen
		"\u00a0": " ",   // non-breaking space
		// Add more replacements as needed
	}

	// Apply all replacements
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	// Trim leading/trailing whitespace and return
	return strings.TrimSpace(input)
}

func ParseFlexibleTime(timeStr string) (time.Time, error) {
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
