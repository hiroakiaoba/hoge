package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello, World!!!`))
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
