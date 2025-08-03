package helper

import (
	"encoding/json"
	"os"

	"github.com/parth469/go-web-scraper/internal/model"
)

func SaveToFile(events []model.Event) error {
	jsonData, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		Log.Error("failed to marshal events to JSON: %w", err)
		return err
	}

	if err := os.WriteFile("./output/events.json", jsonData, 0644); err != nil {
		Log.Error("failed to write events.json: %w", err)
		return err
	}

	return nil
}
