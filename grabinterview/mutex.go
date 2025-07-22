package grabinterview

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func increment() {
	mu.Lock() // Lock the mutex before accessing shared data
	count++
	mu.Unlock() // Unlock the mutex after done
}

// slower is read heavy scenario
func MutexTest() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}

// rwmutex
var rwMu sync.RWMutex
var value int

func readValue() int {
	rwMu.RLock()
	defer rwMu.RUnlock()
	return value
}

func writeValue(v int) {
	rwMu.Lock()
	value = v
	rwMu.Unlock()
}

// Writes are more expensive than with sync.Mutex due to the additional overhead of managing read and write locks.
func RWMutexTest() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			readValue()
		}()
	}
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeValue(3)
	}()
	wg.Wait()
	fmt.Println("Final count:", count)
}
