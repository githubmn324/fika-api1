package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Employee struct {
	Employeeid string `json:"employeeid"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Divisionid string `json:"divisionid"`
}

var employees []Employee

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees = append(employees, Employee{Employeeid: "1", Firstname: "John", Lastname: "Doe", Divisionid: "1"})
	employees = append(employees, Employee{Employeeid: "2", Firstname: "Smith", Lastname: "Taylor", Divisionid: "2"})
	fmt.Println("Endopoint Hi: All Employees Endpoint")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/allEmployees", getAllEmployees)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "Cloud Run 1"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
