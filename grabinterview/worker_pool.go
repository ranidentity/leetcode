package grabinterview

import (
	"fmt"
	"sync"
)

func Worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result <- job
	}
}

func WorkerPool() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important: close channel to signal no more jobs
	go func() {
		wg.Wait()
		close(results) // Close results when all workers are done
	}()

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}

}

func WorkerPoolWithContext() {

}
