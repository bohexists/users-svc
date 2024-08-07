package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Привет")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Query().Get("message"))
	})

	http.HandleFunc("/circle", func(w http.ResponseWriter, r *http.Request) {
		radius := r.URL.Query().Get("radius")
		if radius == "" {
			http.Error(w, "Параметр radius не передан", http.StatusBadRequest)
			return
		}

		radiusFloat, err := strconv.ParseFloat(radius, 32)
		if err != nil {
			http.Error(w, "Неверное значение radius", http.StatusBadRequest)
			return
		}

		area := 3.14159 * radiusFloat * radiusFloat
		fmt.Fprintf(w, "Радиус: %.2f, Площадь: %.2f", radiusFloat, area)
	})

	http.ListenAndServe(":80", nil)
}
