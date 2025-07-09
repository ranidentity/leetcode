package grabinterview

import (
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
