package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":9090", nil)
	fmt.Println(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}
