package initializers

import (
	"os"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedisClient() {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to connect to Redis")
	} else {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_INFO, "[INFO] Connected to Redis")
	}

	RedisClient = redis.NewClient(opts)
}
