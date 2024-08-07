package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Привет")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Query().Get("massage"))
	})

	http.ListenAndServe(":80", nil)
}
