package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "one two switch")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
