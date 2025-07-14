package channel

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	mu   sync.RWMutex
	subs map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		subs: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, msg string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}

func (ps *PubSub) Unsubscribe(topic string, ch <-chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i, sub := range ps.subs[topic] {
		if sub == ch {
			// Remove the channel from slice
			ps.subs[topic] = append(ps.subs[topic][:i], ps.subs[topic][i+1:]...)
			close(sub)
			break
		}
	}
}

func PubSubMain() {
	ps := NewPubSub()

	// Subscriber 1 to news
	newsSub1 := ps.Subscribe("news")
	go func() {
		for msg := range newsSub1 {
			fmt.Printf("News Sub1 received: %s\n", msg)
		}
		fmt.Println("News Sub1 closed")
	}()

	// Subscriber 2 to news
	newsSub2 := ps.Subscribe("news")
	go func() {
		for msg := range newsSub2 {
			fmt.Printf("News Sub2 received: %s\n", msg)
		}
		fmt.Println("News Sub2 closed")
	}()

	// Sports subscriber
	sportsSub := ps.Subscribe("sports")
	go func() {
		for msg := range sportsSub {
			fmt.Printf("Sports Sub received: %s\n", msg)
		}
		fmt.Println("Sports Sub closed")
	}()

	// Publisher
	go func() {
		ps.Publish("news", "Breaking: Go 1.20 released!")
		time.Sleep(100 * time.Millisecond)
		ps.Publish("sports", "Team wins championship")
		time.Sleep(100 * time.Millisecond)
		ps.Publish("news", "Update: Critical security patch")
	}()

	time.Sleep(1 * time.Second)
	ps.Unsubscribe("news", newsSub1)
	time.Sleep(100 * time.Millisecond)
	ps.Publish("news", "This won't reach unsubscribed channel")

	time.Sleep(1 * time.Second) // Wait for messages to process
}
