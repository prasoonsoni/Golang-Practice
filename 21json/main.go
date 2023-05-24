package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string   `json:"username"`
	Email    string   `json:"useremail"`
	Password string   `json:"-"` // here - is given so that nobody is able to consume this field
	Posts    []string `json:"userposts,omitempty"` // omitempty ignores empty fields
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson() {
	users := []user{
		{"Prasoon", "prasoon@gmail.com", "123", []string{"1", "2", "3"}},
		{"Aryan", "aryan@gmail.com", "456", nil},
	}

	// package this data as JSON Data

	finalJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}
