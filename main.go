package main

import (
	"fmt"
	"net/http"
	"main/handlers"
)

func main() {

	go handlers.Count()

	http.HandleFunc("/", handlers.HealthyHandler)
	http.HandleFunc("/random_error", handlers.RandomerrorHandler)
	http.HandleFunc("/timeout", handlers.LongresponseHandler)
	http.HandleFunc("/die20sec", handlers.Die20secHandler)
	fmt.Println("Server listening on port 8095")
	http.ListenAndServe(":8095", nil)
}