package main

import (
	"fmt"
	"net/http"
	"main/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HealthyHandler)
	fmt.Println("Server listening on port 8095")
	http.ListenAndServe(":8095", nil)
}