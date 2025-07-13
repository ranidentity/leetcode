package grabinterview

import (
	"fmt"
	"sync"
)

func MySyncMap() {
	var m sync.Map

	// Store values
	m.Store("key1", "value1")
	m.Store("key2", 42)

	// Load values
	if value, ok := m.Load("key1"); ok {
		fmt.Println("key1:", value)
	}

	// Delete a key
	m.Delete("key2")

	// Load or store (atomic check-then-act)
	actual, loaded := m.LoadOrStore("key3", "default")
	fmt.Printf("key3: %v (loaded: %t)\n", actual, loaded)

	// Range over all key-value pairs
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true // return false to stop iteration
	})
}
