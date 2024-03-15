package configs

import (
	"testing"
	"context"
)

func TestConnectDBSingleton(t *testing.T) {
	client1 := ConnectDB()
	client2 := ConnectDB()

	if client1 != client2 {
		t.Error("Singleton instance not maintained")
	}

	// Additional tests to ensure client1 and client2 are functional MongoDB clients
	if err := client1.Ping(context.TODO(), nil); err != nil {
		t.Errorf("Failed to ping client1: %v", err)
	}

	if err := client2.Ping(context.TODO(), nil); err != nil {
		t.Errorf("Failed to ping client2: %v", err)
	}
}

