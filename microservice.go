package main

import (
	"fmt"
	"net/http"
)

// If can not access this go server via browser, before move forward, check firewalld rather than iptables
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8089", nil)
	fmt.Println(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}
