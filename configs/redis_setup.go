package configs

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

const defaultRedisPort = "6379"

type RedisClientBuilder struct {
	addr     string
	password string
}

func NewRedisClientBuilder() *RedisClientBuilder {
	return &RedisClientBuilder{}
}

func (b *RedisClientBuilder) WithAddr(addr string) *RedisClientBuilder {
	b.addr = addr
	return b
}

func (b *RedisClientBuilder) WithPassword(password string) *RedisClientBuilder {
	b.password = password
	return b
}

func (b *RedisClientBuilder) Build() (*redis.Client, error) {
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = defaultRedisPort
	}

	if b.addr == "" {
		b.addr = "localhost:" + redisPort
	}

	if b.password == "" {
		log.Printf("server: redis password is required")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     b.addr,
		Password: b.password,
		DB:       0,
	})

	log.Printf("connect to redis: http://%s/", b.addr)

	return rdb, nil
}
