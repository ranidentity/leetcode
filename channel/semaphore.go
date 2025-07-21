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

func main() {
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
		taskID := i

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
