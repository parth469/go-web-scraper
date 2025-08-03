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

func SavePoster(posters []model.Poster) error {
	const posterFilePath = "./output/Posters.json"

	// Create output directory if it doesn't exist
	if err := os.MkdirAll("./output", 0755); err != nil {
		Log.Error("failed to create output directory: %w", err)
		return err
	}

	var existingPosters []model.Poster

	fileData, err := os.ReadFile(posterFilePath)
	if err == nil && len(fileData) > 0 {
		if err := json.Unmarshal(fileData, &existingPosters); err != nil {
			Log.Error("failed to unmarshal existing posters: %w", err)
			return err
		}
	} else if err != nil && !os.IsNotExist(err) {
		Log.Error("failed to read Posters.json: %w", err)
		return err
	}

	existingPosters = append(existingPosters, posters...)

	jsonData, err := json.MarshalIndent(existingPosters, "", "  ")
	if err != nil {
		Log.Error("failed to marshal posters to JSON: %w", err)
		return err
	}

	if err := os.WriteFile(posterFilePath, jsonData, 0644); err != nil {
		Log.Error("failed to write Posters.json: %w", err)
		return err
	}

	return nil
}
