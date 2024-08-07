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

	http.HandleFunc("/circle", func(w http.ResponseWriter, r *http.Request) {
		radius := r.URL.Query().Get("radius")
		if radius == "" {
			http.Error(w, "Параметр 'radius' отсутствует", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Радиус: %s", radius)
	})

	http.ListenAndServe(":80", nil)
}
