package api

import (
	"fmt"
	"net/http"
)

func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println("query: ", query)
	message := query["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
