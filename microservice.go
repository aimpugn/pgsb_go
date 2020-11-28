package main

import (
	"fmt"
	"net/http"
	"os"
)

// If can not access this go server via browser, before move forward, check firewalld rather than iptables
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(port(), nil)
	fmt.Println(err)
}

func port() string {
	port := os.Getenv("PORT")
	fmt.Println(port)
	if len(port) == 0 {
		port = "8089"
	}

	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}
