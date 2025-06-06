package main

import (
    "context"
    "encoding/json"
    "log"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/redis/go-redis/v9"
)

var (
    redisClient *redis.Client
    redisTopic  string
)

func init() {
    redisAddr := os.Getenv("REDIS_HOST")
    redisPassword := os.Getenv("REDIS_PASSWORD")
    redisTopic = os.Getenv("REDIS_TOPIC")

    if redisAddr == "" || redisTopic == "" {
        log.Fatal("Missing REDIS_HOST or REDIS_TOPIC environment variable")
    }

    redisClient = redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword,
        DB:       0,
    })
}

func handler(ctx context.Context, event json.RawMessage) error {
    return redisClient.Publish(ctx, redisTopic, event).Err()
}

func main() {
    lambda.Start(handler)
}
