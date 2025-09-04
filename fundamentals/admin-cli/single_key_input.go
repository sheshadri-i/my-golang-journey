// Reference: Learn Go for Beginners Crash Course (Udemy) by Trevor Sawler.
//
// github.com/eiannone/keyboard: This is a popular third-party library that simplifies
// reading keyboard input in Go. It abstracts away the complexity of handling raw
// terminal modes across different operating systems.
//
// You call keyboard.Open() to start listening for key presses, and then use a simple API
// to read events. This package is easy to use and handles cross-platform differences for you.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	serviceRunning := false
	fmt.Println("Server Admin CLI:")
	fmt.Println("Press")
	fmt.Println("  y - Start service")
	fmt.Println("  n - Stop service")
	fmt.Println("  q - Quit")
	fmt.Println("  any other key - Show service status.")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		handleKeyPress(char, key, &serviceRunning)
	}
}

func handleKeyPress(char rune, key keyboard.Key, serviceRunning *bool) {
	if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
		fmt.Println("Quitting admin CLI...")
		os.Exit(0)
	} else if char == 'y' || char == 'Y' {
		if !*serviceRunning {
			*serviceRunning = true
			fmt.Println("Service started.")
		} else {
			fmt.Println("Service is already running.")
		}
	} else if char == 'n' || char == 'N' {
		if *serviceRunning {
			*serviceRunning = false
			fmt.Println("Service stopped.")
		} else {
			fmt.Println("Service is not running.")
		}
	} else {
		status := "stopped"
		if *serviceRunning {
			status = "running"
		}
		fmt.Printf("Status: Service is currently %s.\n", status)
	}
}
