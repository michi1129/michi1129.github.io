package main

import (
	"fmt"
	"net/http"
)

func slash(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	fmt.Fprintf(w, "hello %v", text)
}

func main() {
	http.HandleFunc("/slash", slash)
	http.ListenAndServe(":8090", nil)
}
