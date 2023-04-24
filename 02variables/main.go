package main

import "fmt"

// if first letter of declaration is capital then it is a public variable
const LoginToken string = "scdscsc"

func main() {
	var username string = "Prasoon Soni"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45656723476
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "prasoon.codes"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style for variables using walrus operator(:=) This is only allowed inside a method
	//In Go, := is for declaration + assignment, whereas = is for assignment only.
	// For example, var foo int = 10 is the same as foo := 10.
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)
	
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)


}
