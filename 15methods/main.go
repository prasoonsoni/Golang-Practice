package main

import "fmt"

// Always use capital letter to make it exportable
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Method
func (user User) GetStatus() {
	fmt.Println("Is User Active?", user.Status)
}

// when normally as "user User" this only passes the copy we have to pass by reference
func (user User) ChangeEmail(email string) {
	user.Email = email
	fmt.Println(user)
}

// when functions go in classes they are called as methods, in golang there are no classes so we will take help of structs
func main() {
	fmt.Println("Welcome to Methods")

	prasoon := User{"Prasoon", "me@prasoon.codes", true, 21}
	fmt.Println(prasoon)
	prasoon.GetStatus()

	// After changing email
	prasoon.ChangeEmail("prason@gmail.com")

	fmt.Println(prasoon)
}
