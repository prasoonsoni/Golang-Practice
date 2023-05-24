package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to My Modules")
	greeter()
	// creating new router using mux
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	// logging if any error
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
