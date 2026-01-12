package main

import (
	"fmt"
	"os"
)

// --- Interface ---
type Logger interface {
	Log(message string)
}

// --- Implementations ---

// 1. Console Logger working with struct
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("[CONSOLE]:", msg)
}

// 2. File Logger
type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(msg string) {
	// Simple append to file
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error logging to file:", err)
		return
	}
	defer file.Close()
	file.WriteString(msg + "\n")
	fmt.Println("[FILE_LOGGED]:", msg)
}

// 3. Mock Logger (For Testing - only stores in memory)
type MockLogger struct {
	Logs []string
}

// Note: Using pointer receiver to modify the slice!
func (m *MockLogger) Log(msg string) {
	m.Logs = append(m.Logs, msg)
}

// --- Application ---
type Server struct {
	logger Logger
}

func (s *Server) Start() {
	s.logger.Log("Server starting...")
	// ... logic ...
	s.logger.Log("Server stopped.")
}

func main() {
	// 1. Production Mode
	// srv := Server{logger: FileLogger{Filename: "app.log"}}

	// 2. Local Dev Mode
	srv := Server{logger: ConsoleLogger{}}
	srv.Start()

	// 3. Test Mode (showing verification)
	fmt.Println("--- Testing ---")
	mock := &MockLogger{} // Must be pointer because Log has pointer receiver
	testSrv := Server{logger: mock}
	testSrv.Start()

	if len(mock.Logs) == 2 {
		fmt.Println("Test Passed: Logger received 2 messages")
		fmt.Println("Captured:", mock.Logs)
	} else {
		fmt.Println("Test Failed!")
	}
}
