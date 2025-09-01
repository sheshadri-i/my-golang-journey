// Reference: Mastering Go by Mihalis Tsoukalos
//
// In Go, visibility is controlled by the case of the first letter of
// the variable, function, or method name.
// 1. Public (Exported): If the name starts with a capital letter, you can access it
// 	from other packages. This is ideal for creating the public-facing API of a package.
// 2. Private (Unexported): If the name starts with a lowercase letter, you can only access
// 	it within the same package. Private variables and functions help you encapsulate
// 	internal state and logic, protecting your data and forcing controlled access.

// This file demonstrates how you can access public items from the visibility_demo package,
// and why private items are hidden. This is a key concept for writing robust, maintainable Go code.

package main

import (
	"fmt"
	visibility_demo "my-golang-journey/fundamentals/variables/visibility"
)

func main() {
	fmt.Println("Demonstrating Variable Visibility in Go")

	// You can access the public variable DefaultUser
	fmt.Println("Public Variable:", visibility_demo.DefaultUser)

	// You can call the public function NewUser
	user := visibility_demo.NewUser("Alice", 25)
	fmt.Println("User info:", user.GetInfo())

	// You cannot directly access private fields like 'name' and 'age'.
	// This will cause a compile-time error: 'user.name undefined'
	// fmt.Println(user.name)

	// You cannot access private variables or functions from internal.go
	// These will also cause compile-time errors.
	// visibility_demo.defaultUser
	// visibility_demo.createUser("Bob", 30)
}
