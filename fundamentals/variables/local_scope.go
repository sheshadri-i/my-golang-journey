// Reference: https://go.dev/ref/spec#Short_variable_declarations
package main

import (
	"fmt"
)

// Short Declaration operator (:=) for Local Scopes and Simple Assignments
// Why use it?
// - It is a powerful and concise way for you to declare and initialize one or more variables
//   at the point of use.
// - Go automatically infers the type of the variable from the value you're assigning,
//   making your code clean and readable.
// - It limits to a local scope such as within a function or a block. Hence, it helps
//   you keep your code tight and prevents variable names from polluting the outer scope.

// Demonstrates how you can use the short declaration operator (:=) for local variables within a function
func doSomething() (string, error) {
	result := "Success"
	return result, nil
}

func main() {

	// Here, result and err are declared and assigned simultaneously,
	result, err := doSomething()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
