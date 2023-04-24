package main

import (
	"fmt"
	"sort"
)

// This is more powerful than arrays in go, this is used in place of arrays

func main() {
	fmt.Println("Welcome to Slices")

	var fruits = []string{"Apple", "Kiwi"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	// Adding to slices
	fruits = append(fruits, "Mango", "Banana")
	fmt.Println("Fruits: ", fruits)

	// SLicing same as python
	fruits = append(fruits[1:3])
	fmt.Println("Fruits: ", fruits)
	fruits = append(fruits[:3])
	fmt.Println("Fruits: ", fruits)

	// using make() to create slice
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// using sort module to sort
	highScores = append(highScores, 555,666,777)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println("High Scores: ", highScores)

	// how to remove a value from slices based on values
	var courses =[]string{"react", "js", "node", "swift", "python", "ruby"}
	var index int = 2
	fmt.Println("Courses: ", courses)
	fmt.Println("Removing element from index: ", index)
	courses = append(courses[:index], courses[index+1:]... )
	fmt.Println("Courses: ", courses)
}
