package handlers

import (
	"net/http"
	"fmt"
)

func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a healthy handler!")
}