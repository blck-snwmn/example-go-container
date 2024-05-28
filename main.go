package main

import (
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		// return uuid
		ud, _ := uuid.NewV7()
		w.Write([]byte(ud.String()))
	})
	// listen and serve
	http.ListenAndServe(":8080", nil)
}
