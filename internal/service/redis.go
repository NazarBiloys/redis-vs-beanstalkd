package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// PutQueueMessageToRedis puts a message to a Redis queue
func PutQueueMessageToRedis(queueName, connectionType string) error {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     getConnection(connectionType),
		Password: "",
		DB:       0,
	})

	err := client.RPush(ctx, queueName, String(1000)).Err()

	defer client.Close()

	if err != nil {
		return fmt.Errorf("failed to put message to Redis queue: %v", err)
	}

	return nil
}

// ReadQueueMessageFromRedis reads a message from a Redis queue
func ReadQueueMessageFromRedis(queueName, connectionType string) (string, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     getConnection(connectionType),
		Password: "",
		DB:       0,
	})

	defer client.Close()

	message, err := client.LPop(ctx, queueName).Result()

	if err != nil {
		return "", fmt.Errorf("failed to read message from Redis queue: %v", err)
	}

	return message, nil
}

func getConnection(redisConnectionType string) string {
	if redisConnectionType == "AOF" {
		return "redis-aof:6379"
	}

	return "redis-rdb:6379"
}
