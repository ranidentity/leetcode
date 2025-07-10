package grabinterview

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RunChannelTesting() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4)
	// 4 channel will standby waiting
	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()

}
func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel not ok")
			return
		}
		fmt.Printf("Reader %d Received %d\n", id, val)
	}
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
func dowork(id int, resultch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	resultch <- fmt.Sprintf("id: %d ch: %d ", id, rand.Intn(10))

}
func ChannelTestClosing() {
	start := time.Now()
	resultch := make(chan string)
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go dowork(i, resultch, &wg)
	}
	go func() {
		wg.Wait()
		close(resultch)
	}()
	for res := range resultch {
		fmt.Println(res)
	}
	fmt.Printf("work took %v seconds \n", time.Since(start))
}
