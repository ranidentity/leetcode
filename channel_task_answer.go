package main

import (
	"fmt"
	"leetcode/proto"
	"sync"
)

func Producer(id int, numTasks int, taskChannel chan<- proto.Task, producerWg *sync.WaitGroup) {
	defer producerWg.Done()
	for i := 0; i < numTasks; i++ {
		task := proto.Task{
			ID:      id*1000 + i,
			Payload: fmt.Sprintf("Data task %d - %d", id, i),
		}
		taskChannel <- task
	}
}
func Consumer(id int, taskChannel <-chan proto.Task, consumerWg *sync.WaitGroup) {
	defer consumerWg.Done()
	for task := range taskChannel {
		fmt.Printf("cosnumerID %d task ID %d Payload %s\n", id, task.ID, task.Payload)
	}
}

func ProducerConsumerSystem() {
	taskChannel := make(chan proto.Task, 5)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	numProducer := 3
	tasksPerProducer := 5
	numConsumer := 2

	producerWg.Add(numProducer)
	for i := 1; i <= numProducer; i++ {
		go Producer(i, tasksPerProducer, taskChannel, &producerWg)
	}

	consumerWg.Add(numConsumer)
	for i := 1; i <= numConsumer; i++ {
		go Consumer(i, taskChannel, &consumerWg)
	}

	go func() {
		producerWg.Wait()
		close(taskChannel)
	}()

	consumerWg.Wait()
}
