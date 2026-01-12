package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

// Alert represents a security event sent by an agent
type Alert struct {
	AgentID   string `json:"agent_id"`
	EventType string `json:"event_type"` // e.g., "PROCESS_START", "FILE_MODIFIED"
	Details   string `json:"details"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("/audit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var alert Alert
		if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
			logger.Error("Failed to decode alert", "error", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Simulate "Analysis"
		logger.Info("Security Alert Received",
			"agent", alert.AgentID,
			"type", alert.EventType,
			"details", alert.Details,
		)

		if alert.EventType == "UNAUTHORIZED_ACCESS" {
			logger.Warn("Crypto-miner signature detected!", "agent", alert.AgentID)
		}

		w.WriteHeader(http.StatusOK)
	})

	logger.Info("XDR Server listening on :9090")
	http.ListenAndServe(":9090", nil)
}
