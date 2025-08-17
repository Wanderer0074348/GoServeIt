package main

import (
	"fmt"
	"net/http"
)

func sendToRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello you have requested: ", r.URL.Path)
}

func main() {
	http.HandleFunc("/", sendToRoot)
	http.HandleFunc("/hello", sendToRoot)

	http.ListenAndServe(":8000", nil)
}
