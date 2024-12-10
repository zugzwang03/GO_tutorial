package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zugzwang03/mongodb/router"
)

func main() {
	fmt.Println("Welcome to mongodb")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}