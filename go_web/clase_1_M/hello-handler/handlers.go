package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola\n")
}

func main() {
	http.HandleFunc("/hola", helloHandler)
	http.ListenAndServe(":8080", nil)
}
