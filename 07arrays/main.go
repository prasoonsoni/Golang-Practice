package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Tomato"
	fruits[3] = "Kiwi"

	fmt.Println("Fruit List is: ", fruits)
	fmt.Println("Fruit List Length is: ", len(fruits))

	var vegetables = [5]string{"potato", "mushroom", "peas"}
	fmt.Println("Vegetables List is: ", vegetables)
	fmt.Println("Vegetables List Length is: ", len(vegetables))
}
