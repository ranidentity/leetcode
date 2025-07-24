package datareader

import (
	"context"
	"fmt"
	"sync"
)

func CrawlDataThroughSemaphore() {
	ctx := context.Background()
	apiURLs := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
	}
	sem := make(chan struct{}, len(apiURLs))
	rawChan := make(chan RawItem, 100)
	var wg sync.WaitGroup
	for _, url := range apiURLs {
		url := url
		wg.Add(1)
		sem <- struct{}{}
		go func(url string) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := fetchAPI(ctx, url, rawChan); err != nil {
				fmt.Printf("Failed to fetch %s: %v\n", url, err)
			}
		}(url)
	}
	go func() {
		wg.Wait()
		close(rawChan)
	}()
	var cleaned []CleanedItem
	for raw := range rawChan {
		cleaned = append(cleaned, cleanData(raw))
	}
	for _, item := range cleaned {
		fmt.Printf("Cleaned: %+v\n", item)
	}
}
