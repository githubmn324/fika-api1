// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	log.Print("starting server...")
// 	http.HandleFunc("/", handler)

// 	employees = append(stuffs, Book{Employeeid: "1", Firstname: "John", Lastname: "Doe", Divisionid: "1"})
// 	employees = append(stuffs, Book{Employeeid: "2", Firstname: "Smith", Lastname: "Taylor", Divisionid: ""})
// 	// http.HandleFunc("/api/book/{org_id}", getBooks)

// 	r := http.NewServeMux()
// 	r.HandleFunc("/api/employee/{org_id}", getEmployees).Method("GET")

// 	// Determine port for HTTP service.
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("defaulting to port %s", port)
// 	}

// 	// Start HTTP server.
// 	log.Printf("listening on port %s", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "Cloud Run 1"
// 	}
// 	fmt.Fprintf(w, "Hello %s!\n", name)
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// func getEmployees(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params:=
// 	for _, item := range employees {
// 		if item.Employeeid == params["id"]{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewENcoder(w).Encode(&Employee{})
// }

// // slice
// var employees []Employee

// type Employee struct {
// 	Employeeid  string  `json:"employeeid"`
// 	Firstname 	string `json:"firstname"`
// 	Lastname  	string `json:"lastname"`
// 	Divisonid  	string  `json:"divisionid"`
// }

// // books = append(stuffs, Book{Employeeid: "2", Isbn: "567225", Title: "Book Two", Author: &Author{Firstname: "Smith", Lastname: "Taylor"}})
