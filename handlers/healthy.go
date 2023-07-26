package handlers

import (
	"net/http"
	"fmt"
)

func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a healthy handler!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status 200 - OK")
}