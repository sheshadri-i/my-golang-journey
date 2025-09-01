// Reference: Learning Go by Jon Bodner
// You can declare multiple variables of the same or different types at once using
// a grouped var declaration. This improves code readability by visually grouping
// related variables together.

// This approach is cleaner than declaring each variable on a new line and makes the code's
// intent more obvious. It's a matter of style that contributes to better code organization.

package main

import "fmt"

// Group related configuration variables for clarity
var (
	ServerHost = "localhost"
	ServerPort = 8080
	MaxConnections = 100
)

// You can also use this to declare variables of different types
var (
	UserName string
	LoggedIn bool
)

func main() {
    // The variables are now ready to be used
    fmt.Printf("Server is running on %s:%d\n", ServerHost, ServerPort)
    
    // You can assign values later
    UserName = "admin"
    LoggedIn = true
    fmt.Printf("User %s is logged in: %v\n", UserName, LoggedIn)
}