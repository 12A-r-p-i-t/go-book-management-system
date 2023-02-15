package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/12A-r-p-i-t/bookManagementSystem/routers"
)

func main() {
	fmt.Println("This is a book Management System")
	r := router.Router()
	fmt.Println("Server is getting Started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running on port 4000")
}
