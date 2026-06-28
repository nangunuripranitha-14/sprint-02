package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Creates a local timestamp string matching your system's timezone
	localTimestamp := time.Now().Format("2006-01-02T15:04:05Z")

	// Returning telemetry data object
	response := HealthResponse{
		Status:    "down",
		Service:   "auth-service",
		Version:   "1.0.0",
		Timestamp: localTimestamp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	const PORT = ":5000"
	http.HandleFunc("/health", healthHandler)

	println("Server is running on http://localhost" + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}
