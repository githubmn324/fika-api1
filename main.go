package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	fmt.Println("Endopoint Hit: All Employees Endpoint")
	json.NewEncoder(w).Encode(employees)
}

func main() {

	log.Printf("starting HTTP")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Origin",
			"X-Requested-With",
			"Authorization",
			"Accept",
			"Accept-Encoding",
			"Access-Control-Allow-Credentials",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":              "hello I'm Cloud Run API 1!",
			"authorizationHeaders": c.GetHeader("Authorization"),
		})
	})

	r.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.JSON(http.StatusOK, gin.H{
			"message":              "hello I'm Cloud Run API 1!",
			"authorizationHeaders": c.GetHeader("Authorization"),
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// コメントアウト
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/allEmployees", getAllEmployees)

	// // Determine port for HTTP service.
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("defaulting to port %s", port)
	// }

	// // Start HTTP server.
	// log.Printf("listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "Cloud Run 1"
// 	}
// 	fmt.Fprintf(w, "Hello %s!\n", name)
// }
