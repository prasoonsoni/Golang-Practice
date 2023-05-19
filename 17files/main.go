package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in file"

	// Creating a file
	file, err := os.Create("./demo.txt")
	if err != nil {
		panic(err)
		// shutdown program and shows error
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length:", length)

	// reading a file
	readFile("./demo.txt")
	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
