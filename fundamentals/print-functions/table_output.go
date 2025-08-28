// Demonstrates formatted table output in Go using fmt
package main

import (
	"fmt"
	"strings"
)

// Job represents a background job with a name, status, and retry count.
type Job struct {
	Name    string
	Status  string
	Retries int
}

func main() {
	// Define a slice of jobs
	jobs := []Job{
		Job{Name: "migrate-db", Status: "done", Retries: 0},
		Job{Name: "image-processor", Status: "running", Retries: 2},
		Job{Name: "email-sender", Status: "queued", Retries: 0},
	}

	// Print table header with column names
	fmt.Printf("%-20s %-10s %7s\n", "NAME", "STATUS", "RETRIES")

	// Print a separator line
	fmt.Println(strings.Repeat("-", 40))

	// Print each job's details in table format
	for _, j := range jobs {
		fmt.Printf("%-20s %-10s %7d\n", j.Name, j.Status, j.Retries)
	}
}
