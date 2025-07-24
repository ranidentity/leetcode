package datareader

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type RawItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CleanedItem struct {
	Name  string
	Value int
}

func CrawlDataFromApi() {
	ctx := context.Background()
	apiURLs := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
	}
	rawChan := make(chan RawItem, 100)
	var wg sync.WaitGroup
	for _, url := range apiURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
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

func fetchAPI(ctx context.Context, url string, out chan<- RawItem) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // help to clean up if error

	body, _ := io.ReadAll(resp.Body)
	var items []RawItem
	if err := json.Unmarshal(body, &items); err != nil {
		return err
	}

	for _, item := range items {
		out <- item
	}
	return nil
}

func cleanData(raw RawItem) CleanedItem {
	val := 0
	fmt.Sscanf(raw.Value, "%d", &val)
	return CleanedItem{
		Name:  raw.Name,
		Value: val,
	}
}
