package main

import (
	"github.com/parth469/go-web-scraper/internal/scraper"
	"github.com/parth469/go-web-scraper/utils/config"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func main() {
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
