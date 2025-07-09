package grabinterview

import (
	"fmt"
	"sync"
	"time"
)

const (
	numJobs    = 100
	numWorkers = 5
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep((50 * time.Millisecond))
		result <- job * 2
	}
}

func ProcessJob() {
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
	close(jobs) // when done, this will close
	wg.Wait()
	close(results)
	for res := range results {
		fmt.Println("Result: ", res)
	}
}
