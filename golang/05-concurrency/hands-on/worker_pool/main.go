package main

import (
	"fmt"
	"time"
)

// Job represents work to be done
type Job struct {
	ID    int
	Value int
}

// Result represents the output
type Result struct {
	JobID  int
	Output int
}

// Worker Function
// Consumes jobs from `jobs` channel, sends results to `results`
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j.ID)
		time.Sleep(time.Millisecond * 500) // Simulate expensive work
		results <- Result{JobID: j.ID, Output: j.Value * 2}
		fmt.Printf("Worker %d finished job %d\n", id, j.ID)
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// Buffered channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 1. Start Workers (Pool)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 2. Send Jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs) // No more jobs

	// 3. Collect Results
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("Result: Job %d -> %d\n", res.JobID, res.Output)
	}
}
