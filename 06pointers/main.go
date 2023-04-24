package main

import "fmt"

// Why need pointers
// Sometimes when we are passing the variables to the function its copies are being passes not the real one so they can cause problems while not changing them. We prefer pointers to pass actual value

// Pointer - Direct reference to memory address

func main() {
	fmt.Println("Welcome to Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	// Reference means &
	var ptr = &myNumber
	fmt.Println("Value of actual value is ", ptr)
	fmt.Println("Value of actual value is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is ", myNumber)

}
