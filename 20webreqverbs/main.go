package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request Verbs")
	performGetRequest()
	performPostRequest()
	performPostFormRequest()
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

func performPostRequest() {
	const url string = "http://localhost:8000/post"

	// fake json payload
	requestData := strings.NewReader(`
		{
			"name":"Prasoon Soni",
			"email":"prasoonsoni.work.com"
		}
	`)

	response, err := http.Post(url, "application/json", requestData)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	contentBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)
	fmt.Println(content)
}

func performPostFormRequest() {
	const myurl string = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("name", "Prasoon Soni")
	data.Add("email", "prasoonsoni.work@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
