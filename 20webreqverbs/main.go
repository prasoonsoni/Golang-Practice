package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request Verbs")
	performGetRequest()
}

func performGetRequest() {
	const url string = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// Another method to convert from bytes to string
	var responseString strings.Builder
	byteCount, _ := responseString.Write(bytes)
	fmt.Println(responseString.String())
	// content := string(bytes)
	fmt.Println(byteCount)
}
