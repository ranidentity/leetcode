package workerpool

import (
	"fmt"
	"leetcode/proto"
	"sync"
)

func worker2(workerID int, tasks <-chan proto.FanInTask, result chan<- proto.FanInTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d working on tasks %s \n", workerID, task.ID)
		task.Payload = fmt.Sprintf("job Done by worker %d", workerID)
		result <- task
	}
}

func FanInWorkerPool(tasks ...chan proto.FanInTask) {
	result := make(chan proto.FanInTask)
	jobCh := make(chan proto.FanInTask)

	numWorker := 3
	var wg sync.WaitGroup
	// initiate worker
	for i := range numWorker {
		wg.Add(1)
		// send to worker
		go worker2(i, jobCh, result, &wg)
	}
	// receive the task
	go func() {
		for _, taskCh := range tasks {
			for task := range taskCh {
				jobCh <- task
			}
			// task := <-taskCh
			// jobCh <- task
		}
		close(jobCh)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()
	for i := range result {
		fmt.Printf("Receiving %s Payload: %s\n", i.ID, i.Payload)
	}
}
