package main

import "fmt"

func main() {
	defer fmt.Println("World") // here the line gets cut out
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	// first three then two then one LIFO order
	fmt.Println("Hello")
	myDefer()

	// print order ~> hello, myDefer()[4,3,2,1,0], three, two, one, world
	// the defer function will be placed here and executed
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
