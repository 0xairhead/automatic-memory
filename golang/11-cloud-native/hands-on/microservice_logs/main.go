package main

import (
	"log/slog"
	"os"
)

func main() {
	// 1. Setup JSON Logger (Standard for Cloud/Kibana/Splunk)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Set as default
	slog.SetDefault(logger)

	logger.Info("Starting microservice", "service", "payment-service", "version", "v1.2.3")

	// 2. Simulate Business Logic with Structured Context
	processPayment(logger, "user_123", 99.99)
}

func processPayment(logger *slog.Logger, userID string, amount float64) {
	// Add context to all logs in this function/request scope
	reqLogger := logger.With("user_id", userID, "request_id", "req-abc-123")

	reqLogger.Info("Processing payment", "amount", amount)

	if amount > 1000 {
		reqLogger.Warn("High value transaction detected")
	}

	// Simulate error
	err := simulateDBError()
	if err != nil {
		reqLogger.Error("Database transaction failed", "error", err)
	}
}

func simulateDBError() error {
	return os.ErrPermission
}
