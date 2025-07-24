package datareader

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	numWorkers  = 8
	channelSize = 1000
)

func FileReader() {
	filePath := "large_file.csv" // Change to your CSV file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.ReuseRecord = true // Optional: reuse memory for performance

	// Optionally: read header
	header, err := reader.Read()
	if err == io.EOF {
		log.Fatal("Empty CSV file")
	} else if err != nil {
		log.Fatalf("Failed to read header: %v\n", err)
	}
	fmt.Printf("Header: %v\n", header)

	recordChan := make(chan []string, channelSize)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		go csWorker(i, recordChan, &wg)
	}
	// Process CSV rows one at a time
	rowCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error reading line %d: %v\n", rowCount+2, err)
			continue
		}

		// Make a copy to avoid reuse issue when ReuseRecord=true
		rowCopy := make([]string, len(record))
		copy(rowCopy, record)

		recordChan <- rowCopy

		rowCount++
		if rowCount%100000 == 0 {
			fmt.Printf("Processed %d rows...\n", rowCount)
		}
	}
	close(recordChan)
	wg.Wait()
	fmt.Printf("Completed. Total rows processed: %d\n", rowCount)
}

func processRow(record []string) {
	// Example: print first column
	if len(record) > 0 {
		// Replace this with real transformation or aggregation logic
		// fmt.Println(record[0])
	}
}

func csWorker(id int, recordChan <-chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range recordChan {
		processRow(record)
	}
}
