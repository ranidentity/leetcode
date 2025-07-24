package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisAddr   = "localhost:6379" // Assuming Redis is running locally on default port
	channelName = "my_redis_channel"
)

// Publisher function
func publisher(ctx context.Context, client *redis.Client) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			log.Println("Publisher exiting.")
			return
		case <-ticker.C:
			message := fmt.Sprintf("Hello from publisher! Message #%d", i)
			err := client.Publish(ctx, channelName, message).Err()
			if err != nil {
				log.Printf("Error publishing message: %v", err)
				return
			}
			log.Printf("Published: %s", message)
		}
	}
}

// Subscriber function
func subscriber(ctx context.Context, client *redis.Client) {
	pubsub := client.Subscribe(ctx, channelName)
	defer pubsub.Close()

	// Wait for confirmation that subscription is created
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatalf("Error receiving confirmation: %v", err)
	}
	log.Printf("Subscribed to channel: %s", channelName)

	ch := pubsub.Channel()

	for {
		select {
		case <-ctx.Done():
			log.Println("Subscriber exiting.")
			return
		case msg := <-ch:
			log.Printf("Received: %s (Channel: %s)", msg.Payload, msg.Channel)
		}
	}
}

func main() {
	// Create a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Ping Redis to ensure connection is established
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis!")

	// Run publisher in a goroutine
	go publisher(ctx, rdb)

	// Run subscriber in a goroutine
	go subscriber(ctx, rdb)

	// Keep the main goroutine alive for a while to see messages
	log.Println("Press Ctrl+C to exit.")
	<-ctx.Done()                // Block until context is cancelled (e.g., by Ctrl+C)
	time.Sleep(1 * time.Second) // Give goroutines a moment to clean up
}
