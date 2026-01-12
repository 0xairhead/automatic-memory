package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Domain Model
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// In-Memory Store (Thread-Safe)
var (
	tasks  = []Task{{ID: 1, Title: "Learn Go"}, {ID: 2, Title: "Build API"}}
	nextID = 3
	mu     sync.Mutex
)

func main() {
	// Handlers
	http.HandleFunc("/tasks", tasksHandler)

	fmt.Println("API Server running on :8080...")
	fmt.Println("Try: curl http://localhost:8080/tasks")
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Return all tasks
		mu.Lock()
		json.NewEncoder(w).Encode(tasks)
		mu.Unlock()

	case "POST":
		// Create new task
		var t Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		t.ID = nextID
		nextID++
		tasks = append(tasks, t)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
