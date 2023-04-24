package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	// this is the reader to read input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	// comma ok || comma err syntax
	// this work as try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type for rating is %T", input)
}
