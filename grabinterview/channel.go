package grabinterview

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func ChannelTest() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // need to close or else deadlock if receiver keep waiting
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Received ", val)
		}
	}()
	wg.Wait()
}

func ChannelTest2() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("Sending: ", i)
			ch <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Received: ", val)
		}
	}()

	wg.Wait()
}

func BufferedChannelTest() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ChannelMerging(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v // Blocks if no receiver ready
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func AdvanceChannelMerging(ctx context.Context, chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			select {
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case out <- v: // will block until some1 read from out
				case <-ctx.Done(): // prevent leaking
					return
				}
			case <-ctx.Done():
				return
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ErrorChannel() {
	var wg sync.WaitGroup
	errChan := make(chan error, 3) // buffered to avoid goroutine leaks

	tasks := []func() error{
		func() error {
			// Simulate a successful task
			fmt.Println("Task 1: success")
			return nil
		},
		func() error {
			// Simulate a failed task
			return errors.New("Task 2 failed")
		},
		func() error {
			// Simulate another successful task
			fmt.Println("Task 3: success")
			return nil
		},
	}

	// Launch all tasks concurrently
	for _, task := range tasks {
		wg.Add(1)
		go func(fn func() error) {
			defer wg.Done()
			if err := fn(); err != nil {
				errChan <- err
			}
		}(task)
	}

	// Close error channel once all tasks are done
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Handle any errors
	for err := range errChan {
		fmt.Println("Error:", err)
	}
}
