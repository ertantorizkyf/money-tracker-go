package initializers

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedisClient() {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Printf("[ERR] Failed to connect to Redis")
	} else {
		log.Printf("[INFO] Connected to Redis")
	}

	RedisClient = redis.NewClient(opts)
}
