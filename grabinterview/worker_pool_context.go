package grabinterview

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MainWorkerPoolWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i:
				fmt.Printf("Sent job %d\n", i)
			case <-ctx.Done():
				fmt.Println("Job generator stopped by context")
				return
			}
		}
	}()
	WorkerPoolWithContext(ctx, 3, jobs, results)
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}
	fmt.Println("All work completed")

}

func WorkerPoolWithContext(ctx context.Context, numWorker int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for i := range numWorker {
		wg.Add(1)
		go WorkerWithContext(ctx, &wg, jobs, results, i)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
}

func WorkerWithContext(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, id int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			select {
			case <-ctx.Done():
				fmt.Printf("Worker%d context cancelled during work\n", id)
				return
			case <-time.After(100 * time.Millisecond):
				result := job * 2
				select {
				case results <- result:
					fmt.Printf("worker %d processed job %d -> %d \n", id, job, result)
				case <-ctx.Done():
					fmt.Printf("Worker%d context cancelled during send result\n", id)
					return
				}
			}
		case <-ctx.Done():
			fmt.Printf("Worker%d context cancelled during waiting for job\n", id)
		}
	}
}
