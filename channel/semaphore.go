package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Semaphore struct to manage concurrent access
type Semaphore struct {
	permits chan struct{}
}

// NewSemaphore creates a new semaphore with given size
func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		permits: make(chan struct{}, size),
	}
}

// Acquire gets a permit from the semaphore
func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

// Release returns a permit to the semaphore
func (s *Semaphore) Release() {
	<-s.permits
}

func MySemaphore() {
	// Create a semaphore with 3 concurrent permits
	sem := NewSemaphore(3)

	// Create a wait group to wait for all goroutines
	var wg sync.WaitGroup

	// Number of tasks to execute
	numTasks := 10

	// Context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		// taskID := i

		go func(id int) {
			defer wg.Done()

			// Acquire semaphore permit
			sem.Acquire()
			defer sem.Release()

			// Simulate some work
			select {
			case <-time.After(time.Second * time.Duration(1+id%3)):
				fmt.Printf("Task %d completed\n", id)
			case <-ctx.Done():
				fmt.Printf("Task %d cancelled\n", id)
				return
			}
		}(i)
	}

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}

func MiniSemaphore() {
	const (
		totalTasks    = 10 // Total tasks to run
		maxConcurrent = 3  // Max goroutines allowed at once
	)

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent) // Semaphore channel

	for i := 0; i < totalTasks; i++ {
		wg.Add(1)

		// Acquire semaphore (blocks if full)
		sem <- struct{}{}

		go func(taskID int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore when done

			// Simulate work
			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Second * 2)
			fmt.Printf("Task %d completed\n", taskID)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed")
}

func MiniSemaphoreWithTimeout() {
	const (
		totalTasks    = 10              // Total tasks to run
		maxConcurrent = 3               // Max goroutines allowed at once
		timeout       = 3 * time.Second // Timeout duration
	)

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent) // Semaphore channel

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // Ensure the context is canceled when the function exits

	for i := 0; i < totalTasks; i++ {
		wg.Add(1)

		// Acquire semaphore (blocks if full)
		sem <- struct{}{}

		go func(taskID int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore when done

			select {
			case <-ctx.Done():
				// Context was canceled (timeout exceeded)
				fmt.Printf("Task %d canceled due to timeout\n", taskID)
				return
			default:
				// Simulate work
				fmt.Printf("Task %d started\n", taskID)
				time.Sleep(time.Second * 2)
				fmt.Printf("Task %d completed\n", taskID)
			}
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed or timeout exceeded")
}
