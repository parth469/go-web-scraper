package main

import (
	"os"
	"path/filepath"

	"github.com/parth469/go-web-scraper/internal/scraper"
	"github.com/parth469/go-web-scraper/utils/config"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func main() {
	// Delete Posters.json file before starting
	postersFile := filepath.Join("output", "Posters.json")
	if err := os.Remove(postersFile); err != nil && !os.IsNotExist(err) {
		helper.Log.Error("Failed to delete Posters.json file", err)
	} else {
		helper.Log.Info("Posters.json file deleted successfully")
	}

	err := config.Init()
	if err != nil {
		helper.Log.Fatal("Failed to initialize config", err)
		return
	}

	err = scraper.Init()

	if err != nil {
		helper.Log.Fatal("Failed to initialize scraper", err)
		return
	}
}
