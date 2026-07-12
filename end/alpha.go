package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type Health struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func main() {

	// Connect to NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// Create Health event
	health := Health{
		Status:    "failed",
		Service:   "api-gateway",
		Version:   "v1.0.0",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Convert to JSON
	data, err := json.Marshal(health)
	if err != nil {
		panic(err)
	}

	// Publish to Alpha's subject
	err = nc.Publish("infrastructure.health", data)
	if err != nil {
		panic(err)
	}

	// Ensure the message is sent
	err = nc.Flush()
	if err != nil {
		panic(err)
	}

	fmt.Println("Health event published successfully!")
}