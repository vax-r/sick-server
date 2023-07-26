package handlers

import (
	"net/http"
	"fmt"
)

var count int

func Count() {
	for{
		count = count + 1
	}
}

func Die20secHandler(w http.ResponseWriter, r *http.Request) {
	time := count/20
	if ( time % 2 == 1 ) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Status 400 - Bad Request")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, count)
	fmt.Fprintln(w, "Status 200 - OK")
}