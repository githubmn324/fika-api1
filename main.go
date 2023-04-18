package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Printf("starting HTTP")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "hello I'm Cloud Run API 1!",
			"requstHeader": string(requestDump),
			"auth0-token":  c.Request.Header.Get("auth0-token"),
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
