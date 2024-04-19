package configs

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

const defaultRedisPort = "6379"

func NewRedisClient() *redis.Client {
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisPort == "" {
		redisPort = defaultRedisPort
	}
	if redisPassword == "" {
		log.Printf("server: redis password is required")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + redisPort,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb
}
