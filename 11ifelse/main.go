package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watchout"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println("Result: ", result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num:=3; num <10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
