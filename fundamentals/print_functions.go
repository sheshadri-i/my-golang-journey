package main

import "fmt"

// Reference: Mastering Go by Mihalis Tsoukalos

// main demonstrates the differences between fmt.Print, fmt.Println, and fmt.Printf in Go.
func main() {
	num := 123               // integer variable
	str := "456"             // string variable
	msg := "Have a nice day" // another string variable
	floatVal := 3.14159      // float64 variable

	// Use fmt.Println when you want each value printed on a new line
	fmt.Println(num)      // prints: 123\n
	fmt.Println(str)      // prints: 456\n
	fmt.Println(msg)      // prints: Have a nice day\n
	fmt.Println(floatVal) // prints: 3.14159\n

	// Use fmt.Println to print multiple values separated by spaces, ending with a newline.
	fmt.Println(num, str, msg, floatVal) // prints: 123 456 Have a nice day 3.14159\n

	// Use fmt.Print to print multiple values without spaces or newline.
	fmt.Print(num, str, msg, floatVal) // prints: 123456Have a nice day3.14159
	fmt.Println()                      // prints a newline

	// To mimic Println with Print, manually add spaces and a newline.
	fmt.Print(num, " ", str, " ", msg, " ", floatVal, "\n") // prints: 123 456 Have a nice day 3.14159\n

	// Use fmt.Printf when you need custom formatting, such as controlling decimal places, padding, or combining text and variables.
	fmt.Printf("%d - %s - %s - %.3f\n", num, str, msg, floatVal) // prints: 123 - 456 - Have a nice day - 3.142

	// Examples of using different format specifiers (verbs) with fmt.Printf:
	// Use these when you need specific formatting for types, precision, or representation.
	boolVal := true
	ptr := &num

	fmt.Printf("Integer: %d\n", num)        // %d for decimal integer
	fmt.Printf("Float: %.2f\n", floatVal)   // %f for float, .2 for 2 decimal places
	fmt.Printf("Boolean: %t\n", boolVal)    // %t for boolean
	fmt.Printf("String: %s\n", str)         // %s for string
	fmt.Printf("Quoted string: %q\n", str)  // %q for double-quoted string
	fmt.Printf("Hex: %x\n", num)            // %x for hexadecimal
	fmt.Printf("Type: %T\n", floatVal)      // %T for type of the variable
	fmt.Printf("Pointer: %p\n", ptr)        // %p prints the memory address stored in ptr
	fmt.Printf("Pointer: %p\n", &num)       // %p prints the memory address of num directly
	fmt.Printf("Pointer value: %d\n", *ptr) // %d prints the value pointed to by ptr (dereferencing)
}
