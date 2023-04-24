package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in go, no super or parent
	// no classes
	prasoon := User{"Prasoon", "contact@prasoon.codes", true, 21}
	fmt.Println("User: ", prasoon)
	fmt.Printf("Prasoon Details are: %+v\n", prasoon)
	fmt.Println("Name is: ", prasoon.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


