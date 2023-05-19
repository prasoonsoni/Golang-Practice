package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()
	result := add(3, 5)
	fmt.Println("Result is:", result)
	fmt.Println("Result is:", proAdd(1, 2, 3, 4, 5, 6, 7, 8))

	new_total, msg := multiple(1, 2, 3, 4, 5)
	fmt.Println(new_total, msg)
}

// here int is called function signature
func add(num1 int, num2 int) int {
	return num1 + num2
}

// here we do not know how many parameters
// ... is variadic function
func proAdd(values ...int) int {
	var total int = 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

// returning multiple things
func multiple(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi This is multiple function"
}

func greeter() {
	fmt.Println("Namaste!! from GoLang")
}
