package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	// fmt.Println(days)

	// normal for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println()

	// here i will return index
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println()

	// for each loop or enhanced for loop
	for index, day := range days {
		fmt.Println(index, day)
	}

	// kind of a while loop
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 5 {
			break
		}
		if rougueValue == 2 {
			goto lco
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping here")
}
