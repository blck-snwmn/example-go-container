package main

import (
	"net/http"

	"github.com/google/uuid"
)

func main() {
	// def http.handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// return uuid
		ud, _ := uuid.NewV7()
		w.Write([]byte(ud.String()))
	})
	// listen and serve
	http.ListenAndServe(":8080", nil)
}
