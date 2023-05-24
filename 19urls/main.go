package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println("Welcome to URL")
	fmt.Println(myurl)

	// parsing the url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	// Parts of url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=prasoon",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
