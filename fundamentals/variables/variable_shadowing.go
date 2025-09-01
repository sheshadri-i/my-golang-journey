// Reference: Learning Go by Jon Bodner
// Reference: https://medium.com/@shahpershahin/variable-shadowing-in-go-best-practices-to-avoid-confusions-and-bugs-61e03022b54d

// Shadowing occurs when you declare a new variable with the same name as a variable in an
// outer scope. The inner variable "shadows" the outer one, meaning that within its scope,
// any reference to that variable name will access the new, inner variable, not the old, outer one.
// This can lead to subtle and hard-to-find bugs if you're not careful.

package main

import "fmt"

func main() {
	x := 10                      // Outer variable `x`
	fmt.Println("Initial x:", x) // Prints: Initial x: 10

	if true {
		x := 20                    // Inner variable `x` shadows the outer one
		fmt.Println("Inner x:", x) // Prints: Inner x: 20
	}

	fmt.Println("Final x:", x) // Prints: Final x: 10 (The outer `x` is unchanged)

	// How to Avoid Shadowing
	// Avoiding shadowing is often a matter of being explicit.
	// 1. Use the assignment operator = instead of the short declaration operator :=
	// 	when you intend to reassign an existing variable.
	// 2. Be mindful of scopes. If you are inside a new block (e.g., if, for),
	// 	check if the variable you're using already exists in an outer scope.
	// 3. Use static analysis tools. Linters like go vet or golint can often detect potential
	//    shadowing issues and warn you about them.

	// Example:
	y := 10                      // Outer variable `y`
	fmt.Println("Initial y:", y) // Prints: Initial y: 10

	if true {
		y = 20                     // Reassigns the outer variable `y`, does NOT create a new inner variable
		fmt.Println("Inner y:", y) // Prints: Inner y: 20
	}

	fmt.Println("Final y:", y) // Prints: Final y: 20 (The outer `y` is updated)
	// By using '=' instead of ':=', you avoid creating a new variable and update the existing one.
}
