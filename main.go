// Package main provides a Lambda function that forwards events from AWS EventBridge to a Redis pub/sub channel.
// It receives JSON events from EventBridge and publishes them to a configured Redis topic.
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/redis/go-redis/v9"
)

// Global variables for Redis connection and topic name
var (
	redisClient *redis.Client
	redisTopic  string
)

// init initializes the Redis client connection.
// It requires REDIS_HOST and REDIS_TOPIC environment variables to be set,
// and optionally uses REDIS_PASSWORD if provided.
// The function will terminate the program if required environment variables are missing.
func init() {
	redisAddr := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisTopic = os.Getenv("REDIS_TOPIC")

	if redisAddr == "" || redisTopic == "" {
		log.Fatal("Missing REDIS_HOST or REDIS_TOPIC environment variable")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:      redisAddr,
		Password:  redisPassword,
		DB:        0,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
}

// handler is the Lambda function handler that processes EventBridge events.
// It takes a context and the raw JSON event, then publishes the event to the Redis topic.
// Returns an error if the Redis publish operation fails.
func handler(ctx context.Context, event json.RawMessage) error {
	log.Println("Publishing to Redis channel:", redisTopic)
	err := redisClient.Publish(ctx, redisTopic, []byte(event)).Err()
	if err != nil {
		log.Printf("Failed to publish: %v", err)
	}
	return err
}

// main is the entry point for the Lambda function.
// It registers the handler function with the Lambda runtime.
func main() {
	lambda.Start(handler)
}
