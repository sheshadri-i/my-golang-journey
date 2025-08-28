// Demonstrates printing JSON in Go using encoding/json
// Reference: Adapted from Go Programming Cookbook (Over 75 recipes) by Ian Taylor
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents a user with ID, Name, and Email fields.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a User instance
	user := User{1, "Alice", "alice@example.com"}

	// Marshal user to compact JSON
	jsonData1, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	// Print raw Go struct (default formatting, not JSON)
	fmt.Println(user)
	// Print Go struct with field names included (good for logs)
	fmt.Printf("value: %+v\n", user)
	// Print compact JSON output (recommended for JSON data)
	fmt.Printf("JSON output: %s\n", string(jsonData1))

	// Marshal user to indented (pretty) JSON
	jsonData2, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Print indented JSON output
	fmt.Printf("JSON output:\n%s\n", string(jsonData2))
}
