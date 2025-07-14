package channel

import (
	"fmt"
	"sync"
)

// Fan in
func MergingChannel(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			result <- i
		}
	}

	for _, c := range chs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func FanOutWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)

	}
}

func SplittingChannel(ch <-chan int) {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	for i := range numWorkers {
		wg.Add(1)
		go FanOutWorker(i, jobs, &wg)
	}

	for i := range numJobs {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	fmt.Println("done")
}
