package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string   `json:"username"`
	Email    string   `json:"useremail"`
	Password string   `json:"-"`                   // here - is given so that nobody is able to consume this field
	Posts    []string `json:"userposts,omitempty"` // omitempty ignores empty fields
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
	{
                "username": "Prasoon",
                "useremail": "prasoon@gmail.com",
                "userposts": ["1", "2", "3"]
    }
	`)

	var users user
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON Valid")
		json.Unmarshal(jsonData, &users)
		fmt.Printf("%#v\n", users)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value pair

	var myUserData map[string]interface{}
	json.Unmarshal(jsonData, &myUserData) // Passing by reference so it is not copiedA
	fmt.Println(myUserData)

	for k, v := range myUserData {
		fmt.Printf("Key is %v and Value is %v\n", k, v)
	}
}
