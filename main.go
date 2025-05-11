package main

import (
	"log"
	"os"

	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/router"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.ConnectRedisClient()
	initializers.InitializeBloomFilter()
}

func main() {
	r := router.SetupRouter()

	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
