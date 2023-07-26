package handlers

import (
	"net/http"
	"fmt"
	"time"
)

func LongresponseHandler(w http.ResponseWriter, r *http.Request) {
	// this handler is used to make response timeout on purpose
	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status 200 - OK")
}