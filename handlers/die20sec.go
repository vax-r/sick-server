package handlers

import (
	"net/http"
	"fmt"
	"time"
)

var ch chan int = make(chan int, 1)

/*
Count the time of server execution (not accurately)
*/
func Count() {
	ch <- 0
	for {
		cur_val := <- ch
		ch <- (cur_val+1)
		time.Sleep(1 * time.Second)
	}
}

/*
Determine whether the current time is in the available time slot
Every 20 secs is a time slot
Odd time slot number is legal while even isn't
*/
func Die20secHandler(w http.ResponseWriter, r *http.Request) {
	cur_val, ok := <- ch
	live := 0
	if ok {
		slot := cur_val / 20
		live = slot % 2
		ch <- cur_val
	} else {
		fmt.Fprintln(w, "channel not ready")
	}

	if live > 0 {
		fmt.Fprintln(w, "Status 200 - OK")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "Status 400 - Bad Request")
}