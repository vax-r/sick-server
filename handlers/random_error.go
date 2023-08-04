package handlers

import (
	"net/http"
	"fmt"
	"math/rand"
)

func RandomerrorHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(100)
	if ( num < 50 ) {
		fmt.Fprintln(w, "Status 200 - OK")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "Status 400 - Bad Request")
}