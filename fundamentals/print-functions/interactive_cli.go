// Demonstrates usage of Print, Println, and Printf in Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Print: prompt user for input (no newline)
	fmt.Print("Enter username: ")

	// Println: output with automatic newline after input
	var user string
	fmt.Scanln(&user)
	fmt.Println("Welcome,", user)

	// Printf: formatted output for progress indicator (overwrites line)
	for i := 0; i <= 100; i += 10 {
		time.Sleep(80 * time.Millisecond)              // simulate work
		fmt.Printf("\rDownloading: %3d%% complete", i) // formatted output, no newline
	}

	// Println: finish the line with a newline
	fmt.Println()
}
