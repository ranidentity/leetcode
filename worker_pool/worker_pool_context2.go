package workerpool

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Read context closure and receive tasks from taskCh
func worker(ctx context.Context, workerID int, taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskCh:
			if !ok {
				fmt.Printf("Worker %d: Task channel closed. No more tasks. Shutting down...\n", workerID)
				return // Exits the loop and the goroutine
			}
			// Process the task
			fmt.Printf("Worker %d: Processing task %d...\n", workerID, task)
			time.Sleep(time.Duration(rand.Intn(200)+50) * time.Millisecond)
			fmt.Printf("Worker %d: Finished task %d.\n", workerID, task)
		case <-ctx.Done():
			fmt.Println("closure by context")
			return
		}
	}

}

func WorkerPoolContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	taskCh := make(chan int, 5)
	var wg sync.WaitGroup
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, taskCh, &wg)
	}
	// Simulate sending tasks to the workers
	go func() {
		for i := 1; i <= 15; i++ { // Send 15 tasks
			select {
			case <-ctx.Done():
				// If the context is cancelled while sending tasks, stop sending.
				fmt.Println("Main: Context cancelled while sending tasks. Stopping task production.")
				return
			case taskCh <- i:
				fmt.Printf("Main: Sent task %d\n", i)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate delay between sending tasks
			}
		}
		close(taskCh) // Close the channel after all tasks are sent
		fmt.Println("Main: Task channel closed.")
	}()

	// Allow some time for workers to process tasks
	fmt.Println("Main: Running for 2 seconds, then cancelling context...")
	time.Sleep(2 * time.Second) // Wait for 2 seconds

	// Trigger the cancellation
	fmt.Println("Main: Calling cancel() to stop workers.")
	cancel() // This closes ctx.Done() channel, signaling workers to stop

	// Wait for all workers to gracefully shut down
	fmt.Println("Main: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers have shut down. Exiting.")
}
