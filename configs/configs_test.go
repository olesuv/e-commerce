package configs_test

import (
	"context"
	"os"
	"testing"

	"server.go/configs"
)

func TestConnectDBSingleton(t *testing.T) {
	client1 := configs.ConnectDB()
	client2 := configs.ConnectDB()

	if client1 != client2 {
		t.Error("Singleton instance not maintained")
	}

	if err := client1.Ping(context.TODO(), nil); err != nil {
		t.Errorf("Failed to ping client1: %v", err)
	}

	if err := client2.Ping(context.TODO(), nil); err != nil {
		t.Errorf("Failed to ping client2: %v", err)
	}
}

func TestLoadEnv(t *testing.T) {
	err := configs.LoadEnv()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisClientBuilder_Build(t *testing.T) {
	configs.LoadEnv()

	builder := configs.NewRedisClientBuilder()
	builder.WithAddr("localhost:" + os.Getenv("REDIS_PORT")).WithPassword(os.Getenv("REDIS_PASSWORD"))

	client, err := builder.Build()
	if err != nil {
		t.Errorf("Failed to build Redis client: %v", err)
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		t.Errorf("Failed to ping Redis client: %v", err)
	}

	if pong != "PONG" {
		t.Errorf("Unexpected Redis client response: %s", pong)
	}
}
