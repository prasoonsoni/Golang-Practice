package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")
	// map[key_type]value_type
	languages := make(map[string]string)
	languages["PY"] = "Python"
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"

	fmt.Println("Languages: ", languages)
	fmt.Println("PY: ", languages["PY"])
	
	// Deleting value
	delete(languages, "RB")
	fmt.Println("Languages: ", languages)

	// Loops are interesting in golang
	for _, value := range languages {
		fmt.Printf("For key %v, value is %v\n", value)
	}
}
