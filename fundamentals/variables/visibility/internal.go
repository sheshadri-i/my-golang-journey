// This file contains unexported (private) code for the visibility_demo package.
// By starting names with a lowercase letter, you make them accessible only within this package.
package visibility_demo

import (
	"fmt"
)

// A private (unexported) struct. You cannot access 'user' from outside this package.
type user struct {
	name string
	age  int
}

// A private (unexported) variable. You cannot access 'defaultUser' from outside this package.
var defaultUser = user{name: "guest", age: 0}

// A private (unexported) function for internal logic.
// You cannot call 'createUser' from outside this package.
func createUser(name string, age int) user {
	if name == "" {
		fmt.Println("Name is not provided, so using the default user.")
		return defaultUser
	}

	return user{name, age}
}
