package handlers

import (
	"net/http"
	"fmt"
)

func LongresponseHandler(w http.ResponseWriter, r *http.Request) {
	// todo
	fmt.Fprintln(w, "This is a healthy handler!")
}