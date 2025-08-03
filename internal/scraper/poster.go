package scraper

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/parth469/go-web-scraper/internal/model"
	"github.com/parth469/go-web-scraper/utils/helper"
)

func processAbstract(currentPoster *model.Poster, subLink string, subCollector *colly.Collector, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() { 
		<-sem // Release semaphore
	}()

	done := make(chan bool, 1)

	subCollector.OnHTML("div.fl-rich-text", func(div *colly.HTMLElement) {
		var paragraphs []*colly.HTMLElement
		div.ForEach("p", func(index int, p *colly.HTMLElement) {
			paragraphs = append(paragraphs, p)
		})
		count := len(paragraphs)
		if count >= 4 {
			p := paragraphs[count-4]
			currentPoster.Abstract = p.Text
		}
		done <- true
	})

	subCollector.OnError(func(r *colly.Response, e error) {
		helper.Log.Error("error occurred during subCollector request", e)
		done <- false
	})

	err := subCollector.Visit("https://2025.ccneuro.org" + subLink)
	if err != nil {
		helper.Log.Error("failed to visit poster link", err)
		done <- false
	}

	<-done
}

func ProcessPoster(link string, event model.Event, subCollector *colly.Collector) {
	var (
		PosterList []model.Poster
		mu         sync.Mutex
		wg         sync.WaitGroup
		sem        = make(chan struct{}, 5) // Limits to 5 concurrent goroutines
	)

	subCollector.OnHTML("tbody", func(body *colly.HTMLElement) {
		body.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			currentPoster := model.Poster{
				StartTime: event.StartTime,
				EndTime:   event.EndTime,
				PosterId:  helper.CleanString(tr.ChildText("td:nth-child(1) p")),
				Title:     helper.CleanString(tr.ChildText("td:nth-child(2) p")),
				Presenter: helper.CleanString(tr.ChildText("td:nth-child(3) p")),
				TopicArea: helper.CleanString(tr.ChildText("td:nth-child(4) p")),
				Abstract:  "",
			}

			subLink := tr.ChildAttr("td:nth-child(2) p a", "href")
			
			wg.Add(1) // Increment before starting goroutine
			sem <- struct{}{} // Acquire semaphore slot
			
			go func(poster model.Poster, link string) {
				defer wg.Done() // Ensure Done is called exactly once
				
				// Create new collector for this goroutine
				c := subCollector.Clone()
				processAbstract(&poster, link, c, &wg, sem)
				
				mu.Lock()
				PosterList = append(PosterList, poster)
				mu.Unlock()
				
				helper.Log.Info(poster.PosterId)
			}(currentPoster, subLink)
		})
	})

	subCollector.OnError(func(r *colly.Response, e error) {
		helper.Log.Error("error during main collector request", e)
	})

	subCollector.OnScraped(func(r *colly.Response) {
		wg.Wait() // Wait for all goroutines to complete
		close(sem)
		if err := helper.SavePoster(PosterList); err != nil {
			helper.Log.Error("failed to save posters to file", err)
		}
	})

	err := subCollector.Visit(link)
	if err != nil {
		helper.Log.Error("failed to visit poster link", err)
	}
}