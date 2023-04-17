package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// import ("encoding/json")
// type Employee struct {
// 	Employeeid string `json:"employeeid"`
// 	Firstname  string `json:"firstname"`
// 	Lastname   string `json:"lastname"`
// 	Divisionid string `json:"divisionid"`
// }

// var employees []Employee

// func getAllEmployees(w http.ResponseWriter, r *http.Request) {
// 	employees = append(employees, Employee{Employeeid: "1", Firstname: "John", Lastname: "Doe", Divisionid: "1"})
// 	employees = append(employees, Employee{Employeeid: "2", Firstname: "Smith", Lastname: "Taylor", Divisionid: "2"})
// 	fmt.Println("Endopoint Hit: All Employees Endpoint")
// 	json.NewEncoder(w).Encode(employees)
// }

func main() {

	log.Printf("starting HTTP")

	r := gin.Default()

	// CORSはAPI Gatwayで対応済みの為、以下不要。
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		"http://localhost:3000",
	// 		"http://172.20.10.3:3000",
	// 		"https://localhost:3000",
	// 		"https://172.20.10.3:3000",
	// 	},
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 	},
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Headers",
	// 		"Origin",
	// 		"X-Requested-With",
	// 		"Authorization",
	// 		"Accept",
	// 		"Accept-Encoding",
	// 		"Access-Control-Allow-Credentials",
	// 	},
	// 	AllowCredentials: true,
	// 	MaxAge:           24 * time.Hour,
	// }))

	// r.OPTIONS("/", func(c *gin.Context) {
	// c.Header("Access-Control-Expose-Headers", "*")
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.JSON(204, gin.H{
	// 		"message":              "hello I'm Cloud Run API 1!",
	// 		"authorizationHeaders": c.GetHeader("Authorization"),
	// 	})
	// })

	r.GET("/", func(c *gin.Context) {
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "hello I'm Cloud Run API 1!",
			"requstHeader": string(requestDump),
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
