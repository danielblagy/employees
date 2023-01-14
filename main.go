package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var employees []Employee

func welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("/ is requested")
	fmt.Fprintf(w, "Welcome to employees API!")
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	/*employeesJson, err := json.Marshal(employees)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	json.NewEncoder(w).Encode(employees)
}

func main() {
	// init db
	employees = []Employee{
		{Id: 1, FirstName: "Daniil", LastName: "Blagiy", Title: "intern"},
		{Id: 2, FirstName: "Ivan", LastName: "Thespacebiker", Title: "scientist"},
		{Id: 3, FirstName: "Ostap", LastName: "Rodrigez", Title: "bender"},
	}

	fmt.Println("starting employees server locally on port 5000")

	http.HandleFunc("/", welcome)
	http.HandleFunc("/employees", getEmployees)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
