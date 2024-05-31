package main

import (
	"net/http"

	greet "github.com/blck-snwmn/example-go-container/geet"
	"github.com/google/uuid"
)

func main() {
	g := greet.Greeter{}
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.Greet()))
	})

	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		// return uuid
		ud, _ := uuid.NewV7()
		w.Write([]byte(ud.String()))
	})
	http.HandleFunc("hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})
	// listen and serve
	http.ListenAndServe(":8080", nil)
}
