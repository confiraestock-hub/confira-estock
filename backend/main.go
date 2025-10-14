package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API de Estoque Online - Confira-Estock")
	})
	http.ListenAndServe(":8080", nil)
}
