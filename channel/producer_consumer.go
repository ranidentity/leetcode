package channel

import (
	"fmt"
	"sync"
)

type Task struct {
	ID      int
	Payload string
}

func ProducerConsumerSystem() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int)
	workloads := 10
	var tasks []Task

	// producer goroutines
	wg.Add(workloads)
	for i := 0; i < workloads; i++ {
		go func(taskID int) {
			defer wg.Done()
			mu.Lock()
			tasks = append(tasks, Task{ID: taskID})
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	wg.Add(len(tasks))
	for _, i := range tasks {
		go ProcessTask(i, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("receiving: %d\n", i)
	}
}

func ProcessTask(task Task, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- task.ID
}
