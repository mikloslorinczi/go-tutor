package main

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		log.Fatalf("Failed to write HTTP response %s", err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
