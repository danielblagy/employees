package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcome(response http.ResponseWriter, request *http.Request) {
	log.Println("/ is requested")
	fmt.Fprintf(response, "Welcome to employees API!")
}

func main() {
	fmt.Println("starting employees server locally on port 5000")

	http.HandleFunc("/", welcome)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
