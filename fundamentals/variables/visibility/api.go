// This file contains the exported (public) code for the visibility_demo package.
// By starting names with a uppercase letter, you make them accessible outside this package.

package visibility_demo

import (
	"fmt"
)

// A public (exported) struct.
// Note that its fields (`name` and `age`) are still private
type User struct {
	name string
	age  int
}

// A public (exported) variable.
var DefaultUser = User{name: "defaultUser", age: 99}

// NewUser is a public (exported) function
func NewUser(name string, age int) User {
	// Use the private 'createUser' function from internal.go to get an internal 'user' type
	internalUser := createUser(name, age)

	// Convert internalUser (type user) to User using type conversion
	// This is clearer and less error-prone than manually copying fields
	return User(internalUser)
}

// GetInfo is a public (exported) method on the `User` struct.
func (u User) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", u.name, u.age)
}
