package grabinterview

import (
	"context"
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

func channelExample() {
	ch := make(chan string)
	go func() {
		ch <- "data"
	}()
	msg := <-ch
	fmt.Println(msg)
}
